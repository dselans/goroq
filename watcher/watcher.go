package pwatcher

import (
	"fmt"
	"os"
	"time"

	helper "github.com/dselans/goroq/helper"
	fsnotify "github.com/go-fsnotify/fsnotify"
)

type Watcher struct {
	Projects    <-chan string
	RunQueue    chan<- string
	MonitorDirs []string
}

func New(monitorDirs []string, projects <-chan string, runqueue chan<- string) *Watcher {
	watcherObj := &Watcher{}
	watcherObj.Projects = projects
	watcherObj.RunQueue = runqueue
	watcherObj.MonitorDirs = monitorDirs
	return watcherObj
}

// Recursively crawls through a dir and adds all subdirs as new dirs to monitor
func (w *Watcher) addDir(fswatcher *fsnotify.Watcher, dir string) error {
	dirs := helper
}

func (w *Watcher) newWatcher() (*fsnotify.Watcher, error) {
	fswatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	for project := range w.MonitorDirs {
		if err := w.addDir(&fswatcher, project); err != nil {
			return nil, err
		}
	}

	return &fswatcher, nil
}

func (w *Watcher) Run() {
	fmt.Println("Watcher started...")

	fswatcher, err := w.newWatcher()
	if err != nil {
		fmt.Printf("ERROR: Unable to start fs watcher(s). Error: %v\n", err.Error())
		os.Exit(1)
	}

	for {
		time.Sleep(time.Second * 1)
	}
}
