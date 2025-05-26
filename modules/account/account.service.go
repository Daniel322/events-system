package account_module

import (
	"context"
	"events-system/modules/db"
	"fmt"
	"log"
)

type AccountData struct {
	UserId    string `json:"user_id"`
	AccountId string `json:"account_id"`
	Type      string `json:"type"`
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
	query := "UPDATE accounts SET "
	setIndex := 0
	var values []any

	if data.UserId != "" {
		query += "user_id =" + "$" + string(setIndex)
		setIndex++
		values = append(values, data.UserId)
	}
	if data.AccountId != "" {
		query += "account_id =" + "$" + string(setIndex)
		setIndex++
		values = append(values, data.AccountId)
	}
	if data.Type != "" {
		query += "type =" + "$" + string(setIndex)
		setIndex++
		values = append(values, data.Type)
	}

	query += " WHERE id =" + "$" + string(setIndex) + " RETURNING *"

	fmt.Println(query)

	result, err := db.BaseQuery[Account](operationContext, query, values...)
	if err != nil {
		log.Fatal(err)
	}

	return result, err
}

func DeleteAccount(id string, operationContext context.Context) (bool, error) {
	query := "DELETE FROM accounts WHERE id = $1"
	_, err := db.Connection.Exec(operationContext, query, id)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, err
}

func GetAccountByAccountId(account_id string) (int, error) {
	query := "SELECT COUNT(*) FROM accounts WHERE account_id = $1"

	result, err := db.BaseQuery[CountData](context.Background(), query, account_id)
	return result.Count, err
}

func GetUserIdByAccountId(account_id string) (string, error) {
	query := "SELECT user_id from accounts WHERE account_id = $1"

	result, err := db.BaseQuery[UserIdData](context.Background(), query, account_id)
	return result.UserId, err
}

func GetAccounts(options AccountData) (*[]Account, error) {
	query := "SELECT * FROM accounts"
	var values []any

	// TODO: need to add filter support

	rows, err := db.Connection.Query(context.Background(), query, values...)

	var result []Account

	for rows.Next() {
		var iterationScanValue Account
		err = rows.Scan(
			&iterationScanValue.Id,
			&iterationScanValue.UserId,
			&iterationScanValue.AccountId,
			&iterationScanValue.Type,
			&iterationScanValue.CreatedAt,
			&iterationScanValue.UpdatedAt,
		)
		if err != nil {
			log.Fatal(err)
		} else {
			result = append(
				result,
				Account{
					Id:        string(iterationScanValue.Id),
					UserId:    iterationScanValue.UserId,
					AccountId: iterationScanValue.AccountId,
					Type:      iterationScanValue.Type,
					CreatedAt: iterationScanValue.CreatedAt,
					UpdatedAt: iterationScanValue.UpdatedAt,
				},
			)
		}
	}

	return &result, err
}
