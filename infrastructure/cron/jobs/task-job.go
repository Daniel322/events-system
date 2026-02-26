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
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	// tasksList, err := job.Cron.TasksListAction.Run(ctx)

	// if err != nil {
	// 	return utils.GenerateError("TaskJob", err.Error())
	// }

	// for _, task := range *tasksList {
	// 	state, err := job.Cron.ExecTaskCmd.Validate(commands.ExecTaskData{Id: task.ID})
	// 	if err != nil {
	// 		return utils.GenerateError("TaskJob", err.Error())
	// 	}
	// 	msg, err := job.Cron.ExecTaskCmd.Run(ctx, state)
	// 	if err != nil {
	// 		return utils.GenerateError("TaskJob", err.Error())
	// 	}
	// 	// convert msg.ChatId to int64
	// 	intChatId, err := strconv.ParseInt(msg.ChatId, 10, 64)

	// 	job.Cron.Sender.Send(intChatId, msg.Text)
	// }

	return nil
}
