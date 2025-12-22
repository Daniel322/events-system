package usecases

import (
	"events-system/internal/dto"
	"events-system/pkg/utils"
	"log"
)

func (usecase *InternalUseCases) ExecTask(taskId string) (*dto.InfoAboutTaskForTgProvider, error) {
	taskFindOptions := make(map[string]interface{})
	taskFindOptions["id"] = taskId

	currentTask, err := usecase.TaskService.FindOne(taskFindOptions)

	if err != nil {
		return nil, utils.GenerateError("ExecTask", err.Error())
	}

	eventFindOptions := make(map[string]interface{})
	eventFindOptions["id"] = currentTask.EventId.String()

	currentEvent, err := usecase.EventService.FindOne(eventFindOptions)

	if err != nil {
		return nil, utils.GenerateError("ExecTask", err.Error())
	}

	accFindOptions := make(map[string]interface{})
	accFindOptions["id"] = currentTask.AccountId.String()

	currentAcc, err := usecase.AccountService.FindOne(accFindOptions)

	if err != nil {
		return nil, utils.GenerateError("ExecTask", err.Error())
	}

	transaction := usecase.BaseRepository.CreateTransaction()

	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	ok, err := usecase.TaskService.Delete(currentTask.ID.String(), transaction)

	if !ok || err != nil {
		transaction.Rollback()
		if err != nil {
			return nil, utils.GenerateError("ExecTask", err.Error())
		}
		return nil, utils.GenerateError("ExecTask", "task not deleted")
	}

	newTask, err := usecase.TaskService.Create(
		dto.CreateTaskData{
			EventId:   currentEvent.ID,
			AccountId: currentAcc.ID,
			Type:      currentTask.Type,
			Provider:  currentTask.Provider,
			Date:      currentTask.Date.AddDate(1, 0, 0),
		},
		transaction,
	)

	if err != nil {
		return nil, utils.GenerateError("ExecTask", err.Error())
	}

	log.Println("task creted from cron" + newTask.ID.String())

	textMsg := "Attention!" + " For " + currentTask.Type + " in " + currentEvent.Date.Format("01-02") + " will be event " + currentEvent.Info

	if trRes := transaction.Commit(); trRes.Error != nil {
		return nil, utils.GenerateError("ExecTask", trRes.Error.Error())
	}

	return &dto.InfoAboutTaskForTgProvider{
		ChatId: currentAcc.AccountId,
		Text:   textMsg,
	}, nil
}
