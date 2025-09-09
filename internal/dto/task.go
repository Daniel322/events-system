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

type TaskSliceEvent struct {
	Date     time.Time
	Type     string
	Provider string
}

type InfoAboutTaskForTgProvider struct {
	ChatId int64
	Text   string
}
