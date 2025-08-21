package services

import (
	"errors"
	"events-system/infrastructure/providers/db"
	"events-system/interfaces"
	"events-system/internal/domain"
	"events-system/pkg/utils"
	"log"
	"strconv"
	"time"
)

type TaskService struct {
	Name            string
	DB              *db.Database
	taskRepository  interfaces.Repository[domain.Task, domain.CreateTaskData, domain.UpdateTaskData]
	eventRepository interfaces.Repository[domain.Event, domain.CreateEventData, domain.UpdateEventData]
	accRepository   interfaces.Repository[domain.Account, domain.CreateAccountData, domain.UpdateAccountData]
}

type InfoAboutTaskForTgProvider struct {
	ChatId int64
	Text   string
}

func NewTaskService(
	DB *db.Database,
	repo interfaces.Repository[domain.Task, domain.CreateTaskData, domain.UpdateTaskData],
	eventRepo interfaces.Repository[domain.Event, domain.CreateEventData, domain.UpdateEventData],
	accRepository interfaces.Repository[domain.Account, domain.CreateAccountData, domain.UpdateAccountData],
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
	options["date"] = time.Now().Format("2006-01-02")
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

	transaction := ts.DB.CreateTransaction()

	ok, err := ts.taskRepository.Delete(currentTask.ID.String(), transaction)

	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	if !ok || err != nil {
		if err == nil {
			err = errors.New("something went wrong on delete task")
		}
		transaction.Rollback()
		return nil, utils.GenerateError(ts.Name, err.Error())
	}

	newTask, err := ts.taskRepository.Create(domain.CreateTaskData{
		EventId:   currentEvent.ID,
		AccountId: currentAcc.ID,
		Type:      currentTask.Type,
		Provider:  currentTask.Provider,
		Date:      currentTask.Date.AddDate(1, 0, 0),
	}, transaction)

	if err != nil {
		transaction.Rollback()
		return nil, utils.GenerateError(ts.Name, err.Error())
	}

	log.Println("task creted from cron" + newTask.ID.String())

	textMsg := "Attention!" + " For " + currentTask.Type + " in " + currentEvent.Date.Format("01-02") + " will be event " + currentEvent.Info

	transaction.Commit()

	return &InfoAboutTaskForTgProvider{
		ChatId: chatId,
		Text:   textMsg,
	}, nil

}
