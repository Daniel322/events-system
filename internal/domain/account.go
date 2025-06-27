package domain

import (
	"events-system/internal/utils"
	"reflect"
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

type CreateAccountData struct {
	UserId    string
	AccountId string
	Type      string
}

type UpdateAccountData struct {
	AccountId string
	Type      string
}

var SUPPORTED_TYPES = []string{"http", "telegram", "mail"}

const (
	INVALID_TYPE       = "invalid type"
	INVALID_ACCOUNT_ID = "invalid accountId"
	INVALID_USER_ID    = "invalid user id type"
)

func NewAccountFactory() *AccountFactory {
	return &AccountFactory{
		Name: "AccountFactory",
	}
}

func (af *AccountFactory) Create(data CreateAccountData) (*Account, error) {
	var id uuid.UUID = uuid.New()

	parsedUserId, _, err := utils.ParseId(data.UserId)

	if err != nil {
		return nil, utils.GenerateError(af.Name, INVALID_USER_ID)
	}

	if len(data.AccountId) == 0 || len(data.AccountId) > 50 {
		return nil, utils.GenerateError(af.Name, INVALID_ACCOUNT_ID)
	}

	typeContains := slices.Contains(SUPPORTED_TYPES, data.Type)

	if !typeContains {
		return nil, utils.GenerateError(af.Name, INVALID_TYPE)
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

func (af *AccountFactory) Update(acc *Account, data UpdateAccountData) (*Account, error) {
	var reflectData = reflect.ValueOf(data).Elem()
	var fieldsAccount = 0

	if accountId := reflectData.FieldByName("AccountId"); !accountId.IsValid() || len(data.AccountId) > 50 {
		return nil, utils.GenerateError(af.Name, INVALID_ACCOUNT_ID)
	} else {
		acc.AccountId = data.AccountId
		fieldsAccount++
	}

	typeContains := slices.Contains(SUPPORTED_TYPES, data.Type)

	if !typeContains {
		return nil, utils.GenerateError(af.Name, INVALID_TYPE)
	} else {
		acc.Type = data.Type
		fieldsAccount++
	}

	if fieldsAccount > 0 {
		acc.UpdatedAt = time.Now()
	}

	return acc, nil
}
