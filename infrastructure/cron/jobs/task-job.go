package jobs

import (
	"context"
	"events-system/internal/application/commands"
	"events-system/internal/application/queries"
	"events-system/pkg/utils"
	"log"
	"os"
)

type TaskJob struct {
	Logger  *log.Logger
	Query   *queries.TasksList
	Command *commands.Exectask
}

func NewTaskJob(query *queries.TasksList, cmd *commands.Exectask) *TaskJob {
	var logger = log.New(os.Stdout, "TaskJob"+" ", log.LstdFlags)
	return &TaskJob{
		Logger:  logger,
		Query:   query,
		Command: cmd,
	}
}

func (job *TaskJob) Run() error {
	job.Logger.Println("JOB STARTED")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tasksList, err := job.Query.Run(ctx)

	if err != nil {
		return utils.GenerateError("TaskJob", err.Error())
	}

	for _, task := range *tasksList {
		state, err := job.Command.Validate(commands.ExecTaskData{Id: task.ID})
		if err != nil {
			return utils.GenerateError("TaskJob", err.Error())
		}
		msg, err := job.Command.Run(ctx, state)
		if err != nil {
			return utils.GenerateError("TaskJob", err.Error())
		}
		job.Logger.Println(msg)
		// TODO: add tg send msg
	}

	return nil
}
