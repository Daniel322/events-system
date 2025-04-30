package account_module

import (
	"context"
	"events-system/modules/db"
	"fmt"
	"log"
)

type AccountData struct {
	UserId    string      `json:"user_id"`
	AccountId string      `json:"account_id"`
	Type      AccountType `json:"type"`
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

	if data.UserId != "" {
		query += "user_id =" + "$" + string(setIndex)
		setIndex++
	}
	if data.AccountId != "" {
		query += "account_id =" + "$" + string(setIndex)
		setIndex++
	}
	if data.Type.String() != "" {
		query += "type =" + "$" + string(setIndex)
		setIndex++
	}

	query += " WHERE id =" + "$" + string(setIndex) + " RETURNING *"

	fmt.Println(query)

	result, err := db.BaseQuery[Account](context.Background(), query, []any{data}...)
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

// func GetAccounts(options AccountData) (*[]Account, error) {
// 	query := "SELECT * FROM accounts"

// 	if reflect.ValueOf(options).Elem().NumField() != 0 {

// 	}

// 	result, err := db.Connection.Query(context.Background(), query, []any{options}...)

// 	return result, err
// }
