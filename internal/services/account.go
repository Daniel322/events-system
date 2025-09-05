package services

import (
	"events-system/infrastructure/providers/db"
	"events-system/interfaces"
	"events-system/internal/dto"
	entities "events-system/internal/entity"
	"events-system/pkg/utils"
	"time"

	"github.com/google/uuid"
)

type AccountService struct {
	Name       string
	Repository interfaces.Repository[entities.Account]
}

const (
	INVALID_TYPE       = "invalid type"
	INVALID_ACCOUNT_ID = "invalid accountId"
	INVALID_USER_ID    = "invalid user id type"
)

func NewAccountService(repository interfaces.Repository[entities.Account]) *AccountService {
	return &AccountService{
		Name:       "AccountService",
		Repository: repository,
	}
}

func (service *AccountService) checkAccountId(value string) error {
	if len(value) == 0 || len(value) > 50 {
		return utils.GenerateError(service.Name, INVALID_ACCOUNT_ID)
	}

	return nil
}

func (service *AccountService) checkType(value entities.AccountType) error {
	if isValidType := entities.SUPPORTED_ACCOUNT_TYPES[value]; len(isValidType) == 0 {
		return utils.GenerateError(service.Name, INVALID_TYPE)
	}

	return nil
}

func (service *AccountService) Create(
	data dto.CreateAccountData,
	transaction db.DatabaseInstance,
) (*entities.Account, error) {
	var id uuid.UUID = uuid.New()

	if err := uuid.Validate(data.UserId.String()); err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	err := service.checkAccountId(data.AccountId)

	if err != nil {
		return nil, err
	}

	err = service.checkType(data.Type)

	if err != nil {
		return nil, err
	}

	account := &entities.Account{
		ID:        id,
		UserId:    data.UserId,
		AccountId: data.AccountId,
		Type:      entities.SUPPORTED_ACCOUNT_TYPES[data.Type],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	account, err = service.Repository.Save(*account, transaction)

	if err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	return account, nil
}

func (service *AccountService) Find(options map[string]interface{}) (*[]entities.Account, error) {
	results, err := service.Repository.Find(options)

	return results, err
}

func (service *AccountService) Update(
	id string,
	data dto.UpdateAccountData,
	transaction db.DatabaseInstance,
) (*entities.Account, error) {
	findOptions := make(map[string]interface{})
	findOptions["id"] = id

	accounts, err := service.Find(findOptions)

	if err != nil || len(*accounts) == 0 {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	if len(*accounts) == 0 {
		return nil, utils.GenerateError(service.Name, "current acc with id "+id+" not found")
	}

	currentAccount := (*accounts)[0]

	if isInvalidAccountId := service.checkAccountId(data.AccountId); isInvalidAccountId == nil {
		currentAccount.AccountId = data.AccountId
	}
	if isInvalidType := service.checkType(data.Type); isInvalidType == nil {
		currentAccount.Type = entities.SUPPORTED_ACCOUNT_TYPES[data.Type]
	}

	currentAccount.UpdatedAt = time.Now()

	updatedAcc, err := service.Repository.Save(currentAccount, transaction)

	return updatedAcc, err
}

func (service *AccountService) Delete(id string, transaction db.DatabaseInstance) (bool, error) {
	result, err := service.Repository.Destroy(id, transaction)

	return result, err
}

// -------- move to use cases ------

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
