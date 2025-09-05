package dto

import entities "events-system/internal/entity"

type CreateAccountData struct {
	UserId    string
	AccountId string
	Type      entities.AccountType
}

type UpdateAccountData struct {
	AccountId string
	Type      entities.AccountType
}
