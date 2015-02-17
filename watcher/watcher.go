package watcher

import (
	"fmt"

	helper "github.com/dselans/goroq/helper"
	fsnotify "github.com/go-fsnotify/fsnotify"
)

type Watcher struct {
	ProjectName string
	ProjectDir  string
	RunQueue    chan<- string
}

func New(projectName, projectDir string, runqueue chan<- string) *Watcher {
	watcherObj := &Watcher{}
	watcherObj.ProjectName = projectName
	watcherObj.ProjectDir = projectDir
	watcherObj.RunQueue = runqueue
	return watcherObj
}

// Add all subdirs to an existing fsnotify obiject
func (w *Watcher) RecursiveAdd(watcherObj *fsnotify.Watcher, path string) error {
	subdirs := helper.Subfolders(path)

	for _, dir := range subdirs {
		fmt.Println("Adding watcher for dir:", dir)

		if err := watcherObj.Add(dir); err != nil {
			return err
		}
	}

	return nil
}

// Create a new fsnotify object; slurp in all *_test.go from all project subdirs
func (w *Watcher) NewWatcher() (*fsnotify.Watcher, error) {
	watcherObj, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	if err := w.RecursiveAdd(watcherObj, w.ProjectDir); err != nil {
		helper.CustomExit(fmt.Sprintf("Unable to add initial fs watcher for project %v",
			w.ProjectDir), 1)
	}

	return watcherObj, nil
}

func (w *Watcher) Run() {
	fmt.Printf("Watcher started for project %v...\n", w.ProjectName)

	fswatcher, err := w.NewWatcher()
	if err != nil {
		helper.CustomExit(fmt.Sprintf("ERROR: Unable to start fswatcher. Error: %v\n", err.Error()), 1)
	}

	for {
		select {
		case event := <-fswatcher.Events:
			if (event.Op&fsnotify.Create == fsnotify.Create) || (event.Op&fsnotify.Remove == fsnotify.Remove) || (event.Op&fsnotify.Write == fsnotify.Write) || (event.Op&fsnotify.Rename == fsnotify.Rename) {
				// Add to run/test queue as a change has taken place
				fmt.Println("Stuff has happened: ", event)
				w.RunQueue <- w.ProjectDir

				// Make sure to remove a watch for a dir if it gets deleted

				// If dir, roll through and add to existing watcher
				if helper.IsDir(event.Name) {
					if err := w.RecursiveAdd(fswatcher, event.Name); err != nil {
						fmt.Println("ERROR: Tried to add new fswatcher for dir:", event.Name)
					}
				}
			}
		case err := <-fswatcher.Errors:
			if err != nil {
				fmt.Printf("Error while watching project %v. Error: %v\n", w.ProjectName, err)
			}
		}
	}
}
