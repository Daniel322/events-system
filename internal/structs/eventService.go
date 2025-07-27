package structs

import (
	"events-system/internal/domain"
	"time"

	"github.com/google/uuid"
)

type CreateEventData struct {
	AccountId string
	UserId    string
	Info      string
	Date      time.Time
	Providers []byte
}

type Event struct {
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
