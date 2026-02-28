package jobs

import (
	"context"
	"events-system/interfaces"
	"events-system/internal/application/commands"
	"events-system/internal/application/queries"
	"events-system/pkg/utils"
	"log"
	"os"
	"strconv"
)

type TaskJob struct {
	logger *log.Logger
	sender interfaces.Sender
}

func NewTaskJob(sender interfaces.Sender) *TaskJob {
	var logger = log.New(os.Stdout, "TaskJob"+" ", log.LstdFlags)
	return &TaskJob{
		logger: logger,
		sender: sender,
	}
}

func (job *TaskJob) Run() error {
	job.logger.Println("JOB STARTED")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tasksList, err := queries.TasksList.Run(ctx)

	if err != nil {
		return utils.GenerateError("TaskJob", err.Error())
	}

	for _, task := range *tasksList {
		state, err := commands.ExecTask.Validate(commands.ExecTaskData{Id: task.ID})
		if err != nil {
			return utils.GenerateError("TaskJob", err.Error())
		}
		msg, err := commands.ExecTask.Run(ctx, state)
		if err != nil {
			return utils.GenerateError("TaskJob", err.Error())
		}
		// convert msg.ChatId to int64
		intChatId, err := strconv.ParseInt(msg.ChatId, 10, 64)

		job.sender.Send(intChatId, msg.Text)
	}

	return nil
}
