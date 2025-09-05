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

// func (af *AccountFactory) Update(acc *Account, data UpdateAccountData) (*Account, error) {
// 	var reflectData = reflect.ValueOf(&data).Elem()
// 	var fieldsAccount = 0

// 	if accountId := reflectData.FieldByName("AccountId"); !accountId.IsValid() || len(data.AccountId) > 50 {
// 		return nil, utils.GenerateError(af.Name, INVALID_ACCOUNT_ID)
// 	} else {
// 		acc.AccountId = data.AccountId
// 		fieldsAccount++
// 	}

// 	typeContains := slices.Contains(SUPPORTED_TYPES, data.Type)

// 	if !typeContains {
// 		return nil, utils.GenerateError(af.Name, INVALID_TYPE)
// 	} else {
// 		acc.Type = data.Type
// 		fieldsAccount++
// 	}

// 	if fieldsAccount > 0 {
// 		acc.UpdatedAt = time.Now()
// 	}

// 	return acc, nil
// }
