package runner

import (
	"fmt"
	"os"
	"strings"

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
	if err := os.Chdir(r.Project.Dir); err != nil {
		helper.CustomExit(fmt.Sprintf("Unable to Chdir() to project dir '%v'", r.Project.Dir), 1)
	}

	r.Logger.Info.Printf("Runner (%v): Running test on dir: %v\n", r.Project.Name, dir)

	output, err := helper.ExecCmd("go", "test", "./...")
	if err != nil {
		r.PresentOutput(false, output)
		return
	}

	r.PresentOutput(true, output)
}

func (r *Runner) PresentOutput(success bool, output []byte) {
	prefix := "Test Output"

	if !success {
		r.Logger.Warning.Printf("Runner (%v): Got error(s):\n", r.Project.Name)
		r.Logger.Warning.Printf("Runner (%v):\n", r.Project.Name)
		prefix = "Error Output"
	}

	for _, line := range strings.Split(string(output), "\n") {
		if line == "" {
			continue
		}
		r.Logger.Warning.Printf("Runner (%v): [%v] %v\n", r.Project.Name, prefix, line)
	}

}

func (r *Runner) Run() {
	r.Logger.Info.Printf("Runner started for project '%v'...\n", r.Project.Name)

	for {
		dir := <-r.RunQueue
		go r.RunTest(dir)
	}
}
