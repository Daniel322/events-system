package dto

import (
	entities "events-system/internal/entity"

	"github.com/google/uuid"
)

type CreateAccountData struct {
	UserId    uuid.UUID
	AccountId string
	Type      entities.AccountType
}

type UpdateAccountData struct {
	AccountId string
	Type      entities.AccountType
}
