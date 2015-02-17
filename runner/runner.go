package runner

import (
	"log"
)

type Runner struct {
	RunQueue <-chan string
}

func New(runqueue <-chan string) *Runner {
	runnerObj := &Runner{}
	runnerObj.RunQueue = runqueue
	return runnerObj
}

func (r *Runner) RunTest(dir string) {
	log.Println("Running test on dir:", dir)
}

func (r *Runner) Run() {
	log.Println("Runner started...")

	for {
		dir := <-r.RunQueue
		go r.RunTest(dir)
	}
}
