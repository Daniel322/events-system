package dto

import (
	"events-system/internal/domain"
	"time"

	"github.com/google/uuid"
)

type CreateEventDTO struct {
	AccountId string    `json:"account_id" validate:"required"`
	UserId    string    `json:"user_id" validate:"required"`
	Info      string    `json:"info" validate:"required"`
	Date      time.Time `json:"date" validate:"required"`
	Providers []byte    `json:"providers"`
}

type OutputEvent struct {
	ID           uuid.UUID
	UserId       uuid.UUID
	Info         string
	Date         time.Time
	NotifyLevels []byte
	Providers    []byte
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Tasks        []domain.Task
}
