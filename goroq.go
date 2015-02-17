package main

import (
	"log"
	"time"

	config "github.com/dselans/goroq/config"
	helper "github.com/dselans/goroq/helper"
	runner "github.com/dselans/goroq/runner"
	watcher "github.com/dselans/goroq/watcher"
)

const (
	VERSION string = "0.0.1"
)

func main() {
	opts := handleCliArgs()

	projects, err := config.Read(opts.ConfigFile)
	if err != nil {
		helper.CustomExit(err.Error(), 1)
	}

	runqueue := make(chan string, 100)

	// Start test runner goroutine
	runnerObj := runner.New(projects, runqueue)
	go runnerObj.Run()

	// Start fsnotify goroutines
	for _, p := range projects {
		log.Printf("Launching watcher for project %v with dir: %v\n", p.Name, p.Dir)
		watcherObj := watcher.New(p, runqueue)
		go watcherObj.Run()
	}

	for {
		time.Sleep(time.Second * 1)
	}
}
