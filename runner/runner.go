package runner

import (
	config "github.com/dselans/goroq/config"
	golog "github.com/dselans/goroq/golog"
	helper "github.com/dselans/goroq/helper"
)

type Runner struct {
	RunQueue <-chan string
	Project  config.Project
	Logger   *golog.Logger
}

func New(project config.Project, runQueue <-chan string, logger *golog.Logger) *Runner {
	runnerObj := &Runner{}
	runnerObj.RunQueue = runQueue
	runnerObj.Project = project
	runnerObj.Logger = logger
	return runnerObj
}

func (r *Runner) RunTest(dir string) {
	r.Logger.Info.Printf("Runner (%v): Running test on dir: %v\n", r.Project.Name, dir)

	output, err := helper.ExecCmd("go", "test", dir+"/...")
	if err != nil {
		r.Logger.Warning.Printf("Runner (%v): Problems running test in %v. Error: %v\n", r.Project.Name, dir, err)
		return
	}

	r.Logger.Info.Printf("Runner (%v): [Test Output] %v", output)
}

func (r *Runner) Run() {
	r.Logger.Info.Printf("Runner started for project '%v'...\n", r.Project.Name)

	for {
		dir := <-r.RunQueue
		go r.RunTest(dir)
	}
}
