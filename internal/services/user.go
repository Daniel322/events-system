package services

import (
	"events-system/internal/dto"
	entities "events-system/internal/entity"
	repository "events-system/internal/repositories"
	"events-system/internal/utils"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	Name string
}

type UserData struct {
	Username string
}

const USERNAME_CANT_BE_EMPTY_ERR_MSG = "username cant be empty"

func NewUserService() *UserService {
	return &UserService{
		Name: "UserService",
	}
}

func (us *UserService) create(data UserData) (*entities.User, error) {
	var id uuid.UUID = uuid.New()

	if len(data.Username) == 0 {
		return nil, utils.GenerateError(us.Name, USERNAME_CANT_BE_EMPTY_ERR_MSG)
	}

	var user = entities.User{
		ID:        id,
		Username:  data.Username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

func (us *UserService) update(user *entities.User, data UserData) (*entities.User, error) {
	if len(data.Username) == 0 {
		return nil, utils.GenerateError(us.Name, USERNAME_CANT_BE_EMPTY_ERR_MSG)
	}

	user.Username = data.Username
	user.UpdatedAt = time.Now()

	return user, nil
}

func (us UserService) CreateUser(data dto.UserDataDTO) (*dto.OutputUser, error) {
	transaction := repository.CreateTransaction()

	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	user, err := us.create(UserData{Username: data.Username})

	if err != nil {
		transaction.Rollback()
		return nil, utils.GenerateError(us.Name, err.Error())
	}

	user, err = repository.Create[entities.User]("users", *user, transaction)

	if err != nil {
		transaction.Rollback()
		return nil, utils.GenerateError(us.Name, err.Error())
	}

	// TODO: move to service, now no SOLID
	acc, err := repository.Create("accounts", entities.Account{
		ID:        uuid.New(),
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
	user, err := repository.GetById[entities.User]("users", id)

	if err != nil {
		return nil, utils.GenerateError(us.Name, err.Error())
	}

	options := make(map[string]interface{})

	options["user_id"] = user.ID

	accs, err := repository.GetList[entities.Account]("accounts", options)

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
