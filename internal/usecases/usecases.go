package usecases

import (
	"events-system/interfaces"
	"events-system/internal/dto"
	entities "events-system/internal/entity"
	"events-system/pkg/utils"
	"strconv"
)

type InternalUseCases struct {
	BaseRepository interfaces.BaseRepository
	UserService    interfaces.UserService
	AccountService interfaces.AccountService
	EventService   interfaces.EventService
	TaskService    interfaces.TaskService
}

func NewInternalUseCases(
	repository interfaces.BaseRepository,
	user_service interfaces.UserService,
	acc_service interfaces.AccountService,
	event_service interfaces.EventService,
	task_service interfaces.TaskService,
) *InternalUseCases {
	return &InternalUseCases{
		BaseRepository: repository,
		UserService:    user_service,
		AccountService: acc_service,
		EventService:   event_service,
		TaskService:    task_service,
	}
}

func (usecase *InternalUseCases) CreateUser(data dto.CreateUserInput) (*dto.OutputUser, error) {
	transaction := usecase.BaseRepository.CreateTransaction()

	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	user, err := usecase.UserService.Create(data.Username, transaction)

	if err != nil {
		transaction.Rollback()
		return nil, utils.GenerateError("CreateUser", err.Error())
	}

	acc, err := usecase.AccountService.Create(dto.CreateAccountData{
		UserId:    user.ID,
		AccountId: data.AccountId,
		Type:      data.Type,
	}, transaction)

	if err != nil {
		transaction.Rollback()
		return nil, utils.GenerateError("CreateUser", err.Error())
	}

	var accs []entities.Account
	accs = append(accs, *acc)

	if trRes := transaction.Commit(); trRes.Error != nil {
		return nil, utils.GenerateError("CreateUser", trRes.Error.Error())
	}

	return &dto.OutputUser{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Accounts:  accs,
	}, nil
}

func (usecase *InternalUseCases) GetUser(id string) (*dto.OutputUser, error) {
	findOptions := make(map[string]interface{})
	findOptions["id"] = id

	users, err := usecase.UserService.Find(findOptions)

	if err != nil {
		return nil, utils.GenerateError("GetUser", err.Error())
	}

	if len(*users) == 0 {
		return nil, utils.GenerateError("GetUser", "user not found")
	}

	user := (*users)[0]

	options := make(map[string]interface{})

	options["user_id"] = user.ID

	accs, err := usecase.AccountService.Find(options)

	if err != nil {
		return nil, utils.GenerateError("GetUser", err.Error())
	}

	return &dto.OutputUser{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Accounts:  *accs,
	}, nil
}

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
