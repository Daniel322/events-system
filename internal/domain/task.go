package domain

import (
	"events-system/internal/utils"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID
	EventId   uuid.UUID
	AccountId uuid.UUID
	Date      time.Time
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
	return &TaskFactory{
		Name: "TaskFactory",
	}
}

func (tf *TaskFactory) Create(data CreateTaskData) (*Task, error) {
	var reflectData = reflect.ValueOf(data).Elem()

	if eventId := reflectData.FieldByName("EventId"); !eventId.IsValid() {
		return nil, utils.GenerateError(tf.Name, EVENT_ID_IS_REQUIRED)
	}

	if accountId := reflectData.FieldByName("AccountId"); !accountId.IsValid() {
		return nil, utils.GenerateError(tf.Name, ACCOUNT_ID_IS_REQUIRED)
	}

	if date := reflectData.FieldByName("Date"); !date.IsValid() {
		return nil, utils.GenerateError(tf.Name, DATE_IS_REQUIRED)
	}

	var id uuid.UUID = uuid.New()

	task := Task{
		ID:        id,
		EventId:   data.EventId,
		AccountId: data.AccountId,
		Date:      data.Date,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &task, nil
}

func (tf *TaskFactory) Update(task *Task, data UpdateTaskData) (*Task, error) {
	var reflectData = reflect.ValueOf(data).Elem()

	if date := reflectData.FieldByName("Date"); !date.IsValid() {
		return nil, utils.GenerateError(tf.Name, DATE_IS_REQUIRED)
	} else {
		task.Date = data.Date
		task.UpdatedAt = time.Now()
	}

	return task, nil
}
