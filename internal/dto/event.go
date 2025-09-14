package dto

import (
	entities "events-system/internal/entity"
	"time"

	"github.com/google/uuid"
)

type CreateEventDTO struct {
	AccountId uuid.UUID          `json:"account_id" validate:"required"`
	UserId    uuid.UUID          `json:"user_id" validate:"required"`
	Info      string             `json:"info" validate:"required"`
	Date      time.Time          `json:"date" validate:"required"`
	Providers entities.JsonField `json:"providers"`
}

type CreateEventData struct {
	UserId       uuid.UUID
	Info         string
	Date         time.Time
	NotifyLevels entities.JsonField
	Providers    entities.JsonField
}

type UpdateEventData struct {
	Info         string
	Date         time.Time
	NotifyLevels entities.JsonField
	Providers    entities.JsonField
}

type OutputEvent struct {
	ID           uuid.UUID
	UserId       uuid.UUID
	Info         string
	Date         time.Time
	NotifyLevels entities.JsonField
	Providers    entities.JsonField
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Tasks        []entities.Task
}
