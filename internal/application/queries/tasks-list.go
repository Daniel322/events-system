package queries

import (
	"context"
	"events-system/internal/domain/task"
	"events-system/pkg/utils"
	"log"
	"time"
)

type ITasksList struct {
	logger   *log.Logger
	taskRepo *task.TaskRepo
}

var TasksList *ITasksList

func InitTasksList() {
	var logger = log.New(log.Writer(), "TasksList ", log.LstdFlags)

	TasksList = &ITasksList{
		taskRepo: task.Repository,
		logger:   logger,
	}
}

func (this ITasksList) Run(ctx context.Context) (*[]task.Plain, error) {
	options := make(map[string]interface{})
	options["date"] = time.Now().Format("2006-01-02")
	result := new([]task.Plain)
	ctx = context.WithValue(ctx, "tableName", "tasks")
	ctx = context.WithValue(ctx, "ptr", result)

	err := this.taskRepo.Repository.Find(ctx, options)

	if err != nil {
		return nil, utils.GenerateError("TasksList.Run", err.Error())
	}

	return result, nil
}
