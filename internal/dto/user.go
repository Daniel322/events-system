package dto

import (
	entities "events-system/internal/entity"
	"time"

	"github.com/google/uuid"
)

type UserDataDTO struct {
	Username  string `json:"username" validate:"required"`
	Type      string `json:"type" validate:"required,oneof='mail' 'http'"`
	AccountId string `json:"accountId" validate:"required_if=Type mail"`
}

type CreateUserInput struct {
	Username  string
	Type      entities.AccountType
	AccountId string
}

type OutputUser struct {
	ID        uuid.UUID          `json:"id"`
	Username  string             `json:"username"`
	Accounts  []entities.Account `json:"accounts"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}
