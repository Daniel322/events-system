package domain

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id           string
	UserId       uuid.UUID
	Info         string
	Date         time.Time
	NotifyLevels []byte
	Providers    []byte
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
