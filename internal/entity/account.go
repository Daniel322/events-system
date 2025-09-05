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

var supportedAccountTypes = map[AccountType]string{
	Http:     "http",
	Telegram: "telegram",
	Mail:     "mail",
}

func (acc AccountType) String() string {
	return supportedAccountTypes[acc]
}

// func (af *AccountFactory) Create(data CreateAccountData) (*Account, error) {
// 	var id uuid.UUID = uuid.New()

// 	parsedUserId, _, err := utils.ParseId(data.UserId)

// 	if err != nil {
// 		return nil, utils.GenerateError(af.Name, INVALID_USER_ID)
// 	}

// 	if len(data.AccountId) == 0 || len(data.AccountId) > 50 {
// 		return nil, utils.GenerateError(af.Name, INVALID_ACCOUNT_ID)
// 	}

// 	typeContains := slices.Contains(SUPPORTED_TYPES, data.Type)

// 	if !typeContains {
// 		return nil, utils.GenerateError(af.Name, INVALID_TYPE)
// 	}

// 	var account = Account{
// 		ID:        id,
// 		UserId:    parsedUserId,
// 		AccountId: data.AccountId,
// 		Type:      data.Type,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	return &account, nil
// }

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
