package entities

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID     uuid.UUID
	UserId uuid.UUID
	Info   string
	Date   time.Time
	// TODO: think how move that annotation to other struct for keep clean domain entity
	NotifyLevels JsonField `gorm:"type:jsonb"`
	Providers    JsonField `gorm:"type:jsonb"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

var NOTIFY_LEVELS = JsonField{"month", "week", "tomorrow", "today"}
