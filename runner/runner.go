package runner

import (
	"log"

	config "github.com/dselans/goroq/config"
)

type Runner struct {
	RunQueue <-chan string
	Projects []config.Project
}

func New(projects []config.Project, runqueue <-chan string) *Runner {
	runnerObj := &Runner{}
	runnerObj.RunQueue = runqueue
	runnerObj.Projects = projects
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
