package structs

import (
	"events-system/internal/domain"
	"time"

	"github.com/google/uuid"
)

type CreateUserData struct {
	Username  string
	Type      string
	AccountId string
}

type User struct {
	ID        uuid.UUID        `json:"id"`
	Username  string           `json:"username"`
	Accounts  []domain.Account `json:"accounts"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
}
