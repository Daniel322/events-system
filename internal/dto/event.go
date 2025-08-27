package dto

import (
	entities "events-system/internal/entity"
	"time"

	"github.com/google/uuid"
)

type CreateEventDTO struct {
	AccountId string    `json:"account_id" validate:"required"`
	UserId    string    `json:"user_id" validate:"required"`
	Info      string    `json:"info" validate:"required"`
	Date      time.Time `json:"date" validate:"required"`
	Providers []string  `json:"providers"`
}

type OutputEvent struct {
	ID           uuid.UUID
	UserId       uuid.UUID
	Info         string
	Date         time.Time
	NotifyLevels string
	Providers    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Tasks        []entities.Task
}
