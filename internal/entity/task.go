package entities

import (
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

var SUPPORTED_TYPES = []string{"month", "week", "tomorrow", "today"}

var SUPPORTED_PROVIDERS = []string{"mail", "http", "telegram"}
