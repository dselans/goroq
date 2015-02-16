package runner

import (
	"fmt"
	"math/rand"
	"time"
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
	fmt.Printf("routine %v running job %v\n", id, dir)
	time.Sleep(time.Second * 5)
	fmt.Printf("routine %v is finished!\n", id)
}

func (r *Runner) Run() {
	fmt.Println("Runner started...")

	for {
		project := <-r.RunQueue
		go r.runTest(rand.Intn(65535), project)
	}
}
