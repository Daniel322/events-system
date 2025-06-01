package task_module

import (
	"events-system/modules/db"
	"log"
	"time"
)

type CreateTaskData struct {
	EventId   string    `json:"event_id"`
	AccountId string    `json:"account_id"`
	Date      time.Time `json:"date"`
}

func CreateTask(data CreateTaskData) (*Task, error) {
	task := Task{
		EventId:   data.EventId,
		AccountId: data.AccountId,
		Date:      data.Date,
	}

	result := db.Connection.Create(&task)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return &task, nil
}

func GetEventTasks(event_id string) (*[]Task, error) {
	var tasks []Task

	result := db.Connection.Table("tasks").Where("event_id = ?", event_id).Find(&tasks)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return &tasks, nil
}
