package runner

import (
	"fmt"
	"os"

	"github.com/go-fsnotify/fsnotify"
)

type Runner struct {
	RunQueue <-chan string
}

func New(runqueue <-chan string) *Runner {
	runnerObj := &Runner{}
	runnerObj.RunQueue = runqueue
	return runnerObj
}

func (r *Runner) runTest(id int, dir string) {
}

func (r *Runner) Run() {
	fmt.Println("Runner started...")

	for {

	}
}
