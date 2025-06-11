package domain

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id           string    `json:"id"`
	UserId       uuid.UUID `json:"user_id"`
	Info         string    `json:"info"`
	Date         time.Time `json:"date"`
	NotifyLevels []byte    `json:"notify_levels"`
	Providers    []byte    `json:"providers"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
