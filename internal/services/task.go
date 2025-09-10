package services

import (
	"events-system/infrastructure/providers/db"
	"events-system/interfaces"
	"events-system/internal/dto"
	entities "events-system/internal/entity"
	"events-system/pkg/utils"
	"slices"
	"time"

	"github.com/google/uuid"
)

type TaskService struct {
	Name       string
	Repository interfaces.Repository[entities.Task]
}

const (
	DATE_IS_REQUIRED      = "date is required"
	INVALID_TASK_TYPE     = "type is invalid"
	INVALID_TASK_PROVIDER = "provider is invalid"
)

func NewTaskService(repository interfaces.Repository[entities.Task]) *TaskService {
	return &TaskService{
		Name:       "TaskService",
		Repository: repository,
	}
}

func (service *TaskService) checkDate(value time.Time) error {
	if value.IsZero() {
		return utils.GenerateError(service.Name, INVALID_DATE)
	}
	return nil
}

func (service *TaskService) checkContainsOfSupportedvalues(value string, slice []string, err string) error {
	if isContains := slices.Contains(slice, value); !isContains {
		return utils.GenerateError(service.Name, err)
	}

	return nil
}

func (service *TaskService) Find(options map[string]interface{}) (*[]entities.Task, error) {
	results, err := service.Repository.Find(options)

	return results, err
}

func (service *TaskService) FindOne(options map[string]interface{}) (*entities.Task, error) {
	tasks, err := service.Find(options)

	if err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	if len(*tasks) == 0 {
		return nil, utils.GenerateError(service.Name, "current account not found")
	}

	return &(*tasks)[0], nil
}

func (service *TaskService) Delete(id string, transaction db.DatabaseInstance) (bool, error) {
	result, err := service.Repository.Destroy(id, transaction)

	return result, err
}

func (service *TaskService) Create(
	data dto.CreateTaskData,
	transaction db.DatabaseInstance,
) (*entities.Task, error) {
	var id uuid.UUID = uuid.New()

	if err := uuid.Validate(data.AccountId.String()); err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	if err := uuid.Validate(data.EventId.String()); err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	err := service.checkDate(data.Date)

	if err != nil {
		return nil, err
	}

	err = service.checkContainsOfSupportedvalues(data.Type, entities.SUPPORTED_TYPES, INVALID_TASK_TYPE)

	if err != nil {
		return nil, err
	}

	err = service.checkContainsOfSupportedvalues(data.Provider, entities.SUPPORTED_PROVIDERS, INVALID_TASK_PROVIDER)

	if err != nil {
		return nil, err
	}

	task := &entities.Task{
		ID:        id,
		EventId:   data.EventId,
		AccountId: data.AccountId,
		Type:      data.Type,
		Provider:  data.Provider,
		Date:      data.Date,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	task, err = service.Repository.Save(*task, transaction)

	return task, err
}

func (service *TaskService) Update(
	id string,
	date time.Time,
	transaction db.DatabaseInstance,
) (*entities.Task, error) {
	findOptions := make(map[string]interface{})
	findOptions["id"] = id

	currentTask, err := service.FindOne(findOptions)

	if err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	if IsInvalidDate := service.checkDate(date); IsInvalidDate == nil {
		currentTask.Date = date
		currentTask.UpdatedAt = time.Now()

		currentTask, err := service.Repository.Save(*currentTask, transaction)

		return currentTask, err
	}

	return currentTask, nil
}
