package main

import (
	"fmt"
	// runner "github.com/dselans/goroq/runner"
	// watcher "github.com/dselans/goroq/watcher"
)

const (
	VERSION string = "0.0.1"
)

func main() {
	opts := handleCliArgs()
	fmt.Println(opts)
	// runqueue := make(chan string)
	// projects := make(chan string)
	// monitordirs := []string{"one", "two", "three"}

	// runnerObj := runner.New(runqueue)
	// watcherObj := watcher.New(monitordirs, projects, runqueue)

	// go runnerObj.Run()
	// go watcherObj.Run()

	// for {
	// 	fmt.Println("Main program tick")
	// 	time.Sleep(time.Second * 1)
	// }
}
