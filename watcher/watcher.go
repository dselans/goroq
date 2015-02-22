package watcher

import (
	"fmt"

	config "github.com/dselans/goroq/config"
	golog "github.com/dselans/goroq/golog"
	helper "github.com/dselans/goroq/helper"
	fsnotify "github.com/go-fsnotify/fsnotify"
)

type Watcher struct {
	RunQueue    chan<- string
	Project     config.Project
	WatchedDirs []string
	Logger      *golog.Logger
}

func New(project config.Project, runQueue chan<- string, logger *golog.Logger) *Watcher {
	watcherObj := &Watcher{}
	watcherObj.Project = project
	watcherObj.RunQueue = runQueue
	watcherObj.Logger = logger
	return watcherObj
}

// Add all subdirs to an existing fsnotify obiject
func (w *Watcher) RecursiveAdd(watcherObj *fsnotify.Watcher, path string) error {
	subdirs := helper.Subfolders(path)

	for _, dir := range subdirs {
		w.Logger.Info.Printf("Watcher (%v): Adding watcher for dir '%v'", w.Project.Name, dir)

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
		helper.CustomExit(fmt.Sprintf("ERROR: Unable to add initial fs watcher for project %v",
			w.Project.Dir), 1)
	}

	return watcherObj, nil
}

// Check if a given dir is actively watched via fsnotify
// NOTE: Currently no need for this.
func (w *Watcher) IsWatched(watcherObj *fsnotify.Watcher, dir string) bool {
	for _, watchedDir := range w.WatchedDirs {
		if dir == watchedDir {
			return true
		}
	}

	return false
}

func (w *Watcher) Run() {
	w.Logger.Info.Printf("Watcher started for project %v...\n", w.Project.Name)

	fswatcher, err := w.NewWatcher()
	if err != nil {
		helper.CustomExit(fmt.Sprintf("ERROR: Unable to start fswatcher. Error: %v\n", err.Error()), 1)
	}

	for {
		select {
		case event := <-fswatcher.Events:
			if (event.Op&fsnotify.Create == fsnotify.Create) || (event.Op&fsnotify.Remove == fsnotify.Remove) || (event.Op&fsnotify.Write == fsnotify.Write) || (event.Op&fsnotify.Rename == fsnotify.Rename) {
				// Avoid bombing ourselves with logger events
				if event.Name == w.Project.Log {
					continue
				}

				// Add to run/test queue as a change has taken place
				w.RunQueue <- w.Project.Dir

				// If a new dir is created, make sure to add it to watch list
				if event.Op&fsnotify.Create == fsnotify.Create {
					if helper.IsDir(event.Name) {
						if err := w.RecursiveAdd(fswatcher, event.Name); err != nil {
							w.Logger.Error.Printf("Watcher (%v): Tried to add new fswatcher for dir '%v'", w.Project.Name, event.Name)
						}
					}
				}
			}
		case err := <-fswatcher.Errors:
			if err != nil {
				w.Logger.Error.Printf("Watcher (%v): Ran into fsnotify error(s). Error: %v\n", w.Project.Name, err)
			}
		}
	}
}
