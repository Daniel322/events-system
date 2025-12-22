package usecases

import (
	entities "events-system/internal/entity"
	"events-system/pkg/utils"
	"strconv"
)

func (usecase *InternalUseCases) CheckTGAccount(accountId int64) (*entities.Account, error) {
	options := make(map[string]interface{})
	options["account_id"] = strconv.Itoa(int(accountId))
	currentAccounts, err := usecase.AccountService.Find(options)

	if err != nil {
		return nil, utils.GenerateError("CheckTGAccount", err.Error())
	}

	if len(*currentAccounts) == 0 {
		return nil, nil
	}

	return &(*currentAccounts)[0], nil
}
