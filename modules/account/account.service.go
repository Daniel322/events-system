package account_module

import (
	"context"
	"events-system/modules/db"
	"log"

	"github.com/google/uuid"
)

type AccountData struct {
	UserId    uuid.UUID `json:"user_id"`
	AccountId string    `json:"account_id"`
	Type      string    `json:"type"`
}

type CountData struct {
	Count int `json:"count"`
}

type UserIdData struct {
	UserId string `json:"user_id"`
}

func CreateAccount(data AccountData, operationContext context.Context) (*Account, error) {
	account := Account{
		UserId:    &data.UserId,
		AccountId: &data.AccountId,
		Type:      data.Type,
	}

	result := db.Connection.Table("accounts").Create(&account)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return &account, nil
}

func UpdateAccount(id string, data AccountData, operationContext context.Context) (*Account, error) {
	var account Account

	result := db.Connection.Table("accounts").Model(&account).Where("id = ?", id).Updates(data)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return &account, nil
}

func DeleteAccount(id string, operationContext context.Context) (bool, error) {
	result := db.Connection.Table("accounts").Delete(&Account{}, id)

	if result.Error != nil {
		log.Fatal(result.Error)
		return false, result.Error
	}

	return true, nil
}

func GetAccountByAccountId(account_id string) (int, error) {
	var count int64

	result := db.Connection.Table("accounts").Where("account_id = ?", account_id).Count(&count)

	if result.Error != nil {
		log.Fatal(result.Error)
		return 0, result.Error
	}

	return int(count), nil
}

func GetUserIdByAccountId(account_id string) (*uuid.UUID, error) {
	var account Account

	result := db.Connection.Table("accounts").Where("account_id = ?", account_id).First(&account)

	return account.UserId, result.Error
}

func GetAccounts(options AccountData) (*[]Account, error) {
	var accounts []Account

	result := db.Connection.Table("accounts").Where(options).Find(&accounts)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return &accounts, nil
}
