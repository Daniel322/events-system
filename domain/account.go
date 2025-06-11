package domain

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	AccountId string    `json:"account_id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
