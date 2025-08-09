package dto

import (
	"events-system/internal/domain"
	"time"

	"github.com/google/uuid"
)

type UserDataDTO struct {
	Username  string `json:"username" validate:"required"`
	Type      string `json:"type" validate:"required,oneof='mail' 'http'"`
	AccountId string `json:"accountId" validate:"required_if=Type mail"`
}

type OutputUser struct {
	ID        uuid.UUID        `json:"id"`
	Username  string           `json:"username"`
	Accounts  []domain.Account `json:"accounts"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
}
