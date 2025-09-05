package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateTaskData struct {
	EventId   uuid.UUID
	AccountId uuid.UUID
	Date      time.Time
	Type      string
	Provider  string
}

// need in use case
type TaskSliceEvent struct {
	Date time.Time
	Type string
}

// need in use case
type InfoAboutTaskForTgProvider struct {
	ChatId int64
	Text   string
}
