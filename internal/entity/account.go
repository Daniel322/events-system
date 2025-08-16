package entities

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID
	UserId    uuid.UUID
	AccountId string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
