package main

import (
	"log"
	"time"

	runner "github.com/dselans/goroq/runner"
	watcher "github.com/dselans/goroq/watcher"
)

const (
	VERSION string = "0.0.1"
)

func main() {
	runqueue := make(chan string, 100)

	projects := map[string]string{
		"CustomProject1": "/Users/dselans/tests/dir1",
		"CustomProject2": "/Users/dselans/tests/dir2",
	}

	runnerObj := runner.New(runqueue)
	go runnerObj.Run()

	for projectName, projectDir := range projects {
		log.Printf("Launching watcher for project %v with dir: %v\n", projectName, projectDir)

		watcherObj := watcher.New(projectName, projectDir, runqueue)
		go watcherObj.Run()
	}

	for {
		time.Sleep(time.Second * 1)
	}
}
