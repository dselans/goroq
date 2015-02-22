package main

import (
	"fmt"
	"time"

	config "github.com/dselans/goroq/config"
	golog "github.com/dselans/goroq/golog"
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

	for _, p := range projects {
		// Create logger per project
		loggerObj, loggerErr := golog.New(p.Log, opts.QuietMode)
		if loggerErr != nil {
			helper.CustomExit(fmt.Sprintf("Unable to start logger for project"+
				"'%v'. Error: %v", p.Name, loggerErr), 1)
		}

		// Create runqueue for project
		runQueue := make(chan string, 1)

		// Start runner for project
		runnerObj := runner.New(p, runQueue, loggerObj)
		go runnerObj.Run()

		// Start watcher for project
		watcherObj := watcher.New(p, runQueue, loggerObj)
		go watcherObj.Run()
	}

	for {
		time.Sleep(time.Second * 1)
	}
}
