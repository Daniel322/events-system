package task_module

import "time"

type Task struct {
	Id        string    `json:"id"`
	EventId   string    `json:"event_id"`
	AccountId string    `json:"account_id"`
	Date      time.Time `json:"date"`
}
