package services

import (
	entities "events-system/internal/entity"
	dependency_container "events-system/pkg/di"
	repository "events-system/pkg/repository"
	"events-system/pkg/utils"
)

type AccountService struct {
	Name       string
	Repository *repository.Repository[entities.Account]
}

func NewAccountService() error {
	accRepo, err := repository.NewRepository[entities.Account](repository.Accounts)

	if err != nil {
		return utils.GenerateError("AccountService", err.Error())
	}

	dependency_container.Container.Add("accountService", AccountService{
		Name:       "AccountService",
		Repository: accRepo,
	})

	return err
}

// func (af *AccountService) Create(data entities.CreateAccountData, tranasction db.DatabaseInstance) (*entities.Account, error) {
// 	accountFactory, err := dependency_container.Container.Get("accountFactory")

// 	if err != nil {
// 		return nil, utils.GenerateError(af.Name, err.Error())
// 	}

// 	account, err := accountFactory.(interfaces.AccountFactory).Create(data)

// 	if err != nil {
// 		return nil, utils.GenerateError(af.Name, err.Error())
// 	}

// 	resAcc, err := repository.Create(repository.Accounts, *account, tranasction)

// 	if err != nil {
// 		return nil, utils.GenerateError(af.Name, err.Error())
// 	}

// 	return resAcc, nil
// }

// func (as *AccountService) CheckAccount(accountId int64) (*entities.Account, error) {
// 	var options = map[string]interface{}{}
// 	options["account_id"] = strconv.Itoa(int(accountId))
// 	currentAccounts, err := repository.GetList[entities.Account](repository.Accounts, options)

// 	if err != nil {
// 		return nil, utils.GenerateError(as.Name, err.Error())
// 	}

// 	if len(*currentAccounts) == 0 {
// 		return nil, nil
// 	}

// 	return &(*currentAccounts)[0], nil
// }
