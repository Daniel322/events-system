package services

import (
	"events-system/internal/domain"
	"events-system/internal/providers/db"
	"events-system/internal/repositories"
	"events-system/internal/utils"
	"strconv"
	"time"
)

type TaskService struct {
	Name            string
	DB              *db.Database
	taskRepository  repositories.IRepository[domain.Task, domain.CreateTaskData, domain.UpdateTaskData]
	eventRepository repositories.IRepository[domain.Event, domain.CreateEventData, domain.UpdateEventData]
	accRepository   repositories.IRepository[domain.Account, domain.CreateAccountData, domain.UpdateAccountData]
}

type InfoAboutTaskForTgProvider struct {
	ChatId int64
	Text   string
}

func NewTaskService(
	DB *db.Database,
	repo repositories.IRepository[domain.Task, domain.CreateTaskData, domain.UpdateTaskData],
	eventRepo repositories.IRepository[domain.Event, domain.CreateEventData, domain.UpdateEventData],
	accRepository repositories.IRepository[domain.Account, domain.CreateAccountData, domain.UpdateAccountData],
) *TaskService {
	return &TaskService{
		Name:            "TaskService",
		DB:              DB,
		taskRepository:  repo,
		eventRepository: eventRepo,
		accRepository:   accRepository,
	}
}

func (ts *TaskService) GetListOfTodayTasks() (*[]domain.Task, error) {
	var options = make(map[string]interface{})
	options["date"] = time.Now()
	tasks, err := ts.taskRepository.GetList(options)

	if err != nil {
		return nil, utils.GenerateError(ts.Name, err.Error())
	}

	return tasks, nil
}

func (ts *TaskService) ExecTaskAndGenerateNew(taskId string) (*InfoAboutTaskForTgProvider, error) {
	currentTask, err := ts.taskRepository.GetById(taskId)

	if err != nil {
		return nil, utils.GenerateError(ts.Name, err.Error())
	}

	_, strEventId, err := utils.ParseId(currentTask.EventId)

	if err != nil {
		return nil, utils.GenerateError(ts.Name, err.Error())
	}

	_, strAccId, err := utils.ParseId(currentTask.AccountId)

	if err != nil {
		return nil, utils.GenerateError(ts.Name, err.Error())
	}

	currentEvent, err := ts.eventRepository.GetById(strEventId)

	if err != nil {
		return nil, utils.GenerateError(ts.Name, err.Error())
	}

	currentAcc, err := ts.accRepository.GetById(strAccId)

	if err != nil {
		return nil, utils.GenerateError(ts.Name, err.Error())
	}

	chatId, err := strconv.ParseInt(currentAcc.AccountId, 10, 64)

	if err != nil {
		return nil, utils.GenerateError(ts.Name, err.Error())
	}

	textMsg := "Attention!" + " For " + currentTask.Type + " in " + currentEvent.Date.Format("01-02") + " will be event " + currentEvent.Info

	return &InfoAboutTaskForTgProvider{
		ChatId: chatId,
		Text:   textMsg,
	}, nil

}
