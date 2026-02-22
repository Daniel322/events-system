package jobs

import (
	"log"
	"os"
)

type TaskJob struct {
	Logger *log.Logger
}

func NewTaskJob() *TaskJob {
	var logger = log.New(os.Stdout, "TaskJob"+" ", log.LstdFlags)
	return &TaskJob{
		Logger: logger,
	}
}

func (job *TaskJob) Run() error {
	job.Logger.Println("JOB STARTED")
	// get lists of tasks
	// exec task and send msg in loop

	return nil
}
