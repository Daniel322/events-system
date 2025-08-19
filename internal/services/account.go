package services

import (
	"events-system/infrastructure/providers/db"
	entities "events-system/internal/entity"
	dependency_container "events-system/pkg/di"
	repository "events-system/pkg/repository"
	"events-system/pkg/utils"
	"slices"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type AccountService struct {
	Name string
}

type CreateAccountData struct {
	UserId    string
	AccountId string
	Type      string
}

type UpdateAccountData struct {
	AccountId string
	Type      string
}

var SUPPORTED_TYPES = []string{"http", "telegram", "mail"}

const (
	INVALID_TYPE       = "invalid type"
	INVALID_ACCOUNT_ID = "invalid accountId"
	INVALID_USER_ID    = "invalid user id type"
)

func NewAccountService() *AccountService {
	service := &AccountService{
		Name: "AccountService",
	}

	dependency_container.Container.Add("accountService", service)

	return service
}

// TODO: refactor that method, split entity create (make private) and make public save in repo method
func (af *AccountService) Create(data entities.Account, tranasction db.DatabaseInstance) (*entities.Account, error) {
	var id uuid.UUID = uuid.New()

	parsedUserId, _, err := utils.ParseId(data.UserId)

	if err != nil {
		return nil, utils.GenerateError(af.Name, INVALID_USER_ID)
	}

	if len(data.AccountId) == 0 || len(data.AccountId) > 50 {
		return nil, utils.GenerateError(af.Name, INVALID_ACCOUNT_ID)
	}

	typeContains := slices.Contains(SUPPORTED_TYPES, data.Type)

	if !typeContains {
		return nil, utils.GenerateError(af.Name, INVALID_TYPE)
	}

	var account = entities.Account{
		ID:        id,
		UserId:    parsedUserId,
		AccountId: data.AccountId,
		Type:      data.Type,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	resAcc, err := repository.Create(repository.Accounts, account, tranasction)

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
