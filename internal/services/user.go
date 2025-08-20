package services

import (
	"events-system/interfaces"
	"events-system/internal/dto"
	entities "events-system/internal/entity"
	dependency_container "events-system/pkg/di"
	repository "events-system/pkg/repository"
	"events-system/pkg/utils"
)

type UserService struct {
	Name string
}

type UserData struct {
	Username string
}

const USERNAME_CANT_BE_EMPTY_ERR_MSG = "username cant be empty"

func NewUserService() *UserService {
	service := &UserService{
		Name: "UserService",
	}

	dependency_container.Container.Add("userService", service)

	return service
}

func (us UserService) CreateUserWithAccount(data dto.UserDataDTO) (*dto.OutputUser, error) {
	userFactory, err := dependency_container.Container.Get("userFactory")

	if err != nil {
		return nil, utils.GenerateError(us.Name, err.Error())
	}

	accountService, err := dependency_container.Container.Get("accountService")

	if err != nil {
		return nil, utils.GenerateError(us.Name, err.Error())
	}

	transaction := repository.CreateTransaction()

	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	user, err := userFactory.(interfaces.UserFactory).Create(data.Username)

	if err != nil {
		transaction.Rollback()
		return nil, utils.GenerateError(us.Name, err.Error())
	}

	user, err = repository.Create(repository.Users, *user, transaction)

	if err != nil {
		transaction.Rollback()
		return nil, utils.GenerateError(us.Name, err.Error())
	}

	acc, err := accountService.(interfaces.AccountService).Create(entities.Account{
		UserId:    user.ID,
		AccountId: data.AccountId,
		Type:      data.Type,
	}, transaction)

	if err != nil {
		transaction.Rollback()
		return nil, utils.GenerateError(us.Name, err.Error())
	}

	var accs []entities.Account
	accs = append(accs, *acc)

	if trRes := transaction.Commit(); trRes.Error != nil {
		return nil, utils.GenerateError(us.Name, trRes.Error.Error())
	}

	return &dto.OutputUser{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Accounts:  accs,
	}, nil
}

func (us UserService) GetUser(id string) (*dto.OutputUser, error) {
	user, err := repository.GetById[entities.User](repository.Users, id)

	if err != nil {
		return nil, utils.GenerateError(us.Name, err.Error())
	}

	options := make(map[string]interface{})

	options["user_id"] = user.ID

	accs, err := repository.GetList[entities.Account](repository.Accounts, options)

	if err != nil {
		return nil, utils.GenerateError(us.Name, err.Error())
	}

	return &dto.OutputUser{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Accounts:  *accs,
	}, nil
}
