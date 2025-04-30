package account_module

import (
	"context"
	"events-system/modules/db"
	"log"
)

type CreateAccountData struct {
	UserId    string      `json:"user_id"`
	AccountId string      `json:"account_id"`
	Type      AccountType `json:"type"`
}

func CreateAccount(data CreateAccountData) (*Account, error) {
	const query = "INSERT INTO accounts (user_id, account_id, type) VALUES ($1, $2, $3)"
	result, err := db.BaseQuery[Account](context.Background(), query, data.UserId, data.AccountId, data.Type)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}
