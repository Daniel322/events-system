package services

import (
	"events-system/infrastructure/providers/db"
	"events-system/interfaces"
	entities "events-system/internal/entity"
	dependency_container "events-system/pkg/di"
	repository "events-system/pkg/repository"
	"events-system/pkg/utils"
	"strconv"
)

type AccountService struct {
	Name string
}

func NewAccountService() *AccountService {
	service := &AccountService{
		Name: "AccountService",
	}

	dependency_container.Container.Add("accountService", service)

	return service
}

func (af *AccountService) Create(data entities.CreateAccountData, tranasction db.DatabaseInstance) (*entities.Account, error) {
	accountFactory, err := dependency_container.Container.Get("accountFactory")

	if err != nil {
		return nil, utils.GenerateError(af.Name, err.Error())
	}

	account, err := accountFactory.(interfaces.AccountFactory).Create(data)

	if err != nil {
		return nil, utils.GenerateError(af.Name, err.Error())
	}

	resAcc, err := repository.Create(repository.Accounts, *account, tranasction)

	if err != nil {
		return nil, utils.GenerateError(af.Name, err.Error())
	}

	return resAcc, nil
}

func (as *AccountService) CheckAccount(accountId int64) (*entities.Account, error) {
	var options = map[string]interface{}{}
	options["account_id"] = strconv.Itoa(int(accountId))
	currentAccounts, err := repository.GetList[entities.Account](repository.Accounts, options)

	if err != nil {
		return nil, utils.GenerateError(as.Name, err.Error())
	}

	if len(*currentAccounts) == 0 {
		return nil, nil
	}

	return &(*currentAccounts)[0], nil
}
