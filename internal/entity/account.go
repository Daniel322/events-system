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

type AccountType int

const (
	Http AccountType = iota
	Telegram
	Mail
)

var SUPPORTED_ACCOUNT_TYPES = map[AccountType]string{
	Http:     "http",
	Telegram: "telegram",
	Mail:     "mail",
}

func (acc AccountType) String() string {
	return SUPPORTED_ACCOUNT_TYPES[acc]
}
