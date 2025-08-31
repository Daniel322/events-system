package services

import (
	entities "events-system/internal/entity"
	dependency_container "events-system/pkg/di"
	repository "events-system/pkg/repository"
	"events-system/pkg/utils"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	Name       string
	Repository *repository.Repository[entities.User]
}

type UserData struct {
	Username string
}

const USERNAME_CANT_BE_EMPTY_ERR_MSG = "username cant be empty"

func NewUserService() error {
	userRepo, err := repository.NewRepository[entities.User](repository.Users)

	if err != nil {
		return utils.GenerateError("UserService", err.Error())
	}

	dependency_container.Container.Add("userService",
		UserService{
			Name:       "UserService",
			Repository: userRepo,
		})

	return nil
}

func (us *UserService) checkUsername(username string) error {
	if len(username) == 0 {
		return utils.GenerateError(us.Name, USERNAME_CANT_BE_EMPTY_ERR_MSG)
	}

	return nil
}

func (service *UserService) Create(username string) (*entities.User, error) {
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

	user, err = service.Repository.Save(*user, nil)

	if err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	return user, nil
}

func (service *UserService) Find(options map[string]interface{}) {

}

func (service *UserService) Update() {}

func (service *UserService) Delete() {}

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
