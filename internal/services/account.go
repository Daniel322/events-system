package services

import (
	"events-system/internal/domain"
	"events-system/internal/providers/db"
	"events-system/internal/repositories"
	"events-system/internal/utils"
	"strconv"
)

type AccountService struct {
	Name              string
	DB                *db.Database
	accountRepository repositories.IRepository[domain.Account, domain.CreateAccountData, domain.UpdateAccountData]
}

func NewAccountService(
	db *db.Database,
	accountRepository repositories.IRepository[domain.Account, domain.CreateAccountData, domain.UpdateAccountData],
) *AccountService {
	return &AccountService{
		Name:              "AccountService",
		DB:                db,
		accountRepository: accountRepository,
	}
}

func (as *AccountService) CheckAccount(accountId int64) (*domain.Account, error) {
	var options = map[string]interface{}{}
	options["account_id"] = strconv.Itoa(int(accountId))
	currentAccounts, err := as.accountRepository.GetList(options)

	if err != nil {
		return nil, utils.GenerateError(as.Name, err.Error())
	}

	if len(*currentAccounts) == 0 {
		return nil, nil
	}

	return &(*currentAccounts)[0], nil
}
