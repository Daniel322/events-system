package entities

import (
	dependency_container "events-system/pkg/di"
	"events-system/pkg/utils"
	"log"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID
	EventId   uuid.UUID
	AccountId uuid.UUID
	Date      time.Time
	Type      string
	Provider  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TaskFactory struct {
	Name string
}

type CreateTaskData struct {
	EventId   uuid.UUID
	AccountId uuid.UUID
	Date      time.Time
	Type      string
	Provider  string
}

type UpdateTaskData struct {
	Date time.Time
}

const (
	EVENT_ID_IS_REQUIRED   = "event id is required"
	ACCOUNT_ID_IS_REQUIRED = "account id is required"
	DATE_IS_REQUIRED       = "date is required"
)

func NewTaskFactory() *TaskFactory {
	factory := &TaskFactory{
		Name: "TaskFactory",
	}

	dependency_container.Container.Add("taskFactory", factory)

	return factory
}

func (tf *TaskFactory) Create(data CreateTaskData) (*Task, error) {
	log.Println("CREATE TASK DATA", data)
	// TODO: add validation for type and provider
	var reflectData = reflect.ValueOf(&data).Elem()

	if eventId := reflectData.FieldByName("EventId"); eventId.IsZero() {
		return nil, utils.GenerateError(tf.Name, EVENT_ID_IS_REQUIRED)
	}

	if accountId := reflectData.FieldByName("AccountId"); accountId.IsZero() {
		return nil, utils.GenerateError(tf.Name, ACCOUNT_ID_IS_REQUIRED)
	}

	if date := reflectData.FieldByName("Date"); date.IsZero() {
		return nil, utils.GenerateError(tf.Name, DATE_IS_REQUIRED)
	}

	var id uuid.UUID = uuid.New()

	log.Println("Cid", id)

	task := Task{
		ID:        id,
		EventId:   data.EventId,
		AccountId: data.AccountId,
		Date:      data.Date,
		Type:      data.Type,
		Provider:  data.Provider,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	log.Println("task", task)

	return &task, nil
}

func (tf *TaskFactory) Update(task *Task, data UpdateTaskData) (*Task, error) {
	var reflectData = reflect.ValueOf(&data).Elem()

	if date := reflectData.FieldByName("Date"); !date.IsValid() {
		return nil, utils.GenerateError(tf.Name, DATE_IS_REQUIRED)
	} else {
		task.Date = data.Date
		task.UpdatedAt = time.Now()
	}

	return task, nil
}
