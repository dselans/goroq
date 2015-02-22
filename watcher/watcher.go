package watcher

import (
	"fmt"
	"log"

	config "github.com/dselans/goroq/config"
	helper "github.com/dselans/goroq/helper"
	fsnotify "github.com/go-fsnotify/fsnotify"
)

type Watcher struct {
	RunQueue    chan<- string
	Project     config.Project
	WatchedDirs []string
}

func New(project config.Project, runqueue chan<- string) *Watcher {
	watcherObj := &Watcher{}
	watcherObj.Project = project
	watcherObj.RunQueue = runqueue
	return watcherObj
}

// Add all subdirs to an existing fsnotify obiject
func (w *Watcher) RecursiveAdd(watcherObj *fsnotify.Watcher, path string) error {
	subdirs := helper.Subfolders(path)

	for _, dir := range subdirs {
		log.Println("Adding watcher for dir:", dir)

		if err := watcherObj.Add(dir); err != nil {
			return err
		}

		w.WatchedDirs = append(w.WatchedDirs, dir)
	}

	return nil
}

// Create a new fsnotify object; slurp in all *_test.go from all project subdirs
func (w *Watcher) NewWatcher() (*fsnotify.Watcher, error) {
	watcherObj, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	if err := w.RecursiveAdd(watcherObj, w.Project.Dir); err != nil {
		helper.CustomExit(fmt.Sprintf("Unable to add initial fs watcher for project %v",
			w.Project.Dir), 1)
	}

	return watcherObj, nil
}

// Check if a given dir is actively watched via fsnotify
func (w *Watcher) IsWatched(watcherObj *fsnotify.Watcher, dir string) bool {
	for _, watchedDir := range w.WatchedDirs {
		if dir == watchedDir {
			return true
		}
	}

	return false
}

func (w *Watcher) Run() {
	log.Printf("Watcher started for project %v...\n", w.Project.Name)

	fswatcher, err := w.NewWatcher()
	if err != nil {
		helper.CustomExit(fmt.Sprintf("ERROR: Unable to start fswatcher. Error: %v\n", err.Error()), 1)
	}

	for {
		select {
		case event := <-fswatcher.Events:
			if (event.Op&fsnotify.Create == fsnotify.Create) || (event.Op&fsnotify.Remove == fsnotify.Remove) || (event.Op&fsnotify.Write == fsnotify.Write) || (event.Op&fsnotify.Rename == fsnotify.Rename) {
				// Add to run/test queue as a change has taken place
				log.Println("Stuff has happened: ", event)
				w.RunQueue <- w.Project.Dir

				// If a new dir is created, make sure to add it to watch list
				if event.Op&fsnotify.Create == fsnotify.Create {
					if helper.IsDir(event.Name) {
						if err := w.RecursiveAdd(fswatcher, event.Name); err != nil {
							log.Println("ERROR: Tried to add new fswatcher for dir:", event.Name)
						}
					}
				}
			}
		case err := <-fswatcher.Errors:
			if err != nil {
				log.Printf("Error while watching project %v. Error: %v\n", w.Project.Name, err)
			}
		}
	}
}
