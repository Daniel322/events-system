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

func CreateAccount(data AccountData) (*Account, error) {
	const query = "INSERT INTO accounts (user_id, account_id, type) VALUES ($1, $2, $3) RETURNING *"
	result, err := db.BaseQuery[Account](context.Background(), query, data.UserId, data.AccountId, data.Type)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

func UpdateAccount(id string, data AccountData) (*Account, error) {
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

	result, err := db.BaseQuery[Account](context.Background(), query, values...)
	if err != nil {
		log.Fatal(err)
	}

	return result, err
}

func DeleteAccount(id string) (bool, error) {
	query := "DELETE FROM accounts WHERE id = $1"
	_, err := db.Connection.Exec(context.Background(), query, id)
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
