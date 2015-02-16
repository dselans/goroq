package watcher

import (
	"fmt"
	"time"
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

func (w *Watcher) Run() {
	fmt.Println("Watcher started...")

	i := 0

	for {
		time.Sleep(time.Second * 1)
		w.RunQueue <- fmt.Sprintf("Project %v", i)
		w.RunQueue <- fmt.Sprintf("Project %v", i+6521)
		w.RunQueue <- fmt.Sprintf("Project %v")
		i++
	}
}
