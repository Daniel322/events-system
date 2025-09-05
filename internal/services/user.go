package services

import (
	"events-system/infrastructure/providers/db"
	"events-system/interfaces"
	entities "events-system/internal/entity"
	"events-system/pkg/utils"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	Name       string
	Repository interfaces.Repository[entities.User]
}

type UserData struct {
	Username string
}

const USERNAME_CANT_BE_EMPTY_ERR_MSG = "username cant be empty"

func NewUserService(repository interfaces.Repository[entities.User]) *UserService {
	return &UserService{
		Name:       "UserService",
		Repository: repository,
	}
}

func (us *UserService) checkUsername(username string) error {
	if len(username) == 0 {
		return utils.GenerateError(us.Name, USERNAME_CANT_BE_EMPTY_ERR_MSG)
	}

	return nil
}

func (service *UserService) Create(username string, transaction db.DatabaseInstance) (*entities.User, error) {
	var id uuid.UUID = uuid.New()

	err := service.checkUsername(username)

	if err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	user := &entities.User{
		ID:        id,
		Username:  username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err = service.Repository.Save(*user, transaction)

	if err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	return user, nil
}

func (service *UserService) Find(options map[string]interface{}) (*[]entities.User, error) {
	users, err := service.Repository.Find(options)

	return users, err
}

func (service *UserService) Update(id string, username string, transaction db.DatabaseInstance) (*entities.User, error) {
	err := service.checkUsername(username)

	if err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	findOptions := make(map[string]interface{})
	findOptions["id"] = id

	findResult, err := service.Find(findOptions)

	if err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	if len(*findResult) == 0 {
		return nil, utils.GenerateError(service.Name, "current user with id "+id+" not found")
	}

	currentUser := (*findResult)[0]

	currentUser.Username = username
	currentUser.UpdatedAt = time.Now()

	updatedUser, err := service.Repository.Save(currentUser, transaction)

	return updatedUser, err
}

func (service *UserService) Delete(id string, transaction db.DatabaseInstance) (bool, error) {
	result, err := service.Repository.Destroy(id, transaction)

	return result, err
}

// ---------- MOVE THAT CODE TO USE CASES ------------

// func (us UserService) CreateUserWithAccount(data dto.UserDataDTO) (*dto.OutputUser, error) {
// 	userFactory, err := dependency_container.Container.Get("userFactory")

// 	if err != nil {
// 		return nil, utils.GenerateError(us.Name, err.Error())
// 	}

// 	accountService, err := dependency_container.Container.Get("accountService")

// 	if err != nil {
// 		return nil, utils.GenerateError(us.Name, err.Error())
// 	}

// 	transaction := repository.CreateTransaction()

// 	defer func() {
// 		if r := recover(); r != nil {
// 			transaction.Rollback()
// 		}
// 	}()

// 	user, err := userFactory.(interfaces.UserFactory).Create(data.Username)

// 	if err != nil {
// 		transaction.Rollback()
// 		return nil, utils.GenerateError(us.Name, err.Error())
// 	}

// 	user, err = repository.Create(repository.Users, *user, transaction)

// 	if err != nil {
// 		transaction.Rollback()
// 		return nil, utils.GenerateError(us.Name, err.Error())
// 	}

// 	acc, err := accountService.(interfaces.AccountService).Create(entities.CreateAccountData{
// 		UserId:    user.ID.String(),
// 		AccountId: data.AccountId,
// 		Type:      data.Type,
// 	}, transaction)

// 	if err != nil {
// 		transaction.Rollback()
// 		return nil, utils.GenerateError(us.Name, err.Error())
// 	}

// 	var accs []entities.Account
// 	accs = append(accs, *acc)

// 	if trRes := transaction.Commit(); trRes.Error != nil {
// 		return nil, utils.GenerateError(us.Name, trRes.Error.Error())
// 	}

// 	return &dto.OutputUser{
// 		ID:        user.ID,
// 		Username:  user.Username,
// 		CreatedAt: user.CreatedAt,
// 		UpdatedAt: user.UpdatedAt,
// 		Accounts:  accs,
// 	}, nil
// }

// func (us UserService) GetUser(id string) (*dto.OutputUser, error) {
// 	user, err := repository.GetById[entities.User](repository.Users, id)

// 	if err != nil {
// 		return nil, utils.GenerateError(us.Name, err.Error())
// 	}

// 	options := make(map[string]interface{})

// 	options["user_id"] = user.ID

// 	accs, err := repository.GetList[entities.Account](repository.Accounts, options)

// 	if err != nil {
// 		return nil, utils.GenerateError(us.Name, err.Error())
// 	}

// 	return &dto.OutputUser{
// 		ID:        user.ID,
// 		Username:  user.Username,
// 		CreatedAt: user.CreatedAt,
// 		UpdatedAt: user.UpdatedAt,
// 		Accounts:  *accs,
// 	}, nil
// }

// func (us UserService) UpdateUser(id string, username string) (*dto.OutputUser, error) {
// 	_, err := repository.Update(repository.Users, id, entities.User{Username: username}, nil)

// 	if err != nil {
// 		return nil, utils.GenerateError(us.Name, err.Error())
// 	}

// 	currentUser, err := us.GetUser(id)

// 	if err != nil {
// 		return nil, utils.GenerateError(us.Name, err.Error())
// 	}

// 	return currentUser, nil

// }
