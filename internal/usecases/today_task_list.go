package usecases

import (
	entities "events-system/internal/entity"
	"events-system/pkg/utils"
	"time"
)

func (usecase *InternalUseCases) GetListOfTodayTasks() (*[]entities.Task, error) {
	options := make(map[string]interface{})
	options["date"] = time.Now().Format("2006-01-02")

	tasks, err := usecase.TaskService.Find(options)

	if err != nil {
		return nil, utils.GenerateError("GetListOfTodayTasks", err.Error())
	}

	return tasks, nil
}
