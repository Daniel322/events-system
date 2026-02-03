package account

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	value string
	Type  string
}

type Account struct {
	ID        uuid.UUID
	Value     string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
