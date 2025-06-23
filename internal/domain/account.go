package domain

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

type AccountFactory struct {
	Name string
}

type AccountData struct {
	UserId    uuid.UUID
	AccountId string
	Type      string
}

func NewAccountFactory() *AccountFactory {
	return &AccountFactory{
		Name: "accountFactory",
	}
}
