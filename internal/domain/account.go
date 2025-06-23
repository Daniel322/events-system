package domain

import (
	"errors"
	"events-system/internal/utils"
	"fmt"
	"slices"
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
	UserId    string
	AccountId string
	Type      string
}

var SUPPORTED_TYPES = []string{"http", "telegram"}

func NewAccountFactory() *AccountFactory {
	return &AccountFactory{
		Name: "accountFactory",
	}
}

func (af *AccountFactory) CreateAccount(data AccountData) (*Account, error) {
	var id uuid.UUID = uuid.New()

	parsedUserId, _, err := utils.ParseId(data.UserId)

	if err != nil {
		fmt.Println("Invalid userId type")
	}

	if len(data.AccountId) == 0 || len(data.AccountId) > 50 {
		return nil, errors.New("invalid accountId")
	}

	typeContains := slices.Contains(SUPPORTED_TYPES, data.Type)

	if !typeContains {
		return nil, errors.New("invalid type")
	}

	var account = Account{
		ID:        id,
		UserId:    parsedUserId,
		AccountId: data.AccountId,
		Type:      data.Type,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &account, nil
}

func (af *AccountFactory) UpdateAccount(acc *Account, data AccountData) (*Account, error) {
	// TODO: think what fields we can update, after complete func
	acc.UpdatedAt = time.Now()

	return acc, nil
}
