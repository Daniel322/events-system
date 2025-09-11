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

var USERNAME_CANT_BE_EMPTY_ERR_MSG = "username cant be empty"

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

func (service *UserService) FindOne(options map[string]interface{}) (*entities.User, error) {
	users, err := service.Find(options)

	if err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	if len(*users) == 0 {
		return nil, utils.GenerateError(service.Name, "current user not found")
	}

	return &(*users)[0], nil
}

func (service *UserService) Update(
	id string,
	username string,
	transaction db.DatabaseInstance,
) (*entities.User, error) {
	err := service.checkUsername(username)

	if err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	findOptions := make(map[string]interface{})
	findOptions["id"] = id

	currentUser, err := service.FindOne(findOptions)

	if err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	currentUser.Username = username
	currentUser.UpdatedAt = time.Now()

	currentUser, err = service.Repository.Save(*currentUser, transaction)

	return currentUser, err
}

func (service *UserService) Delete(id string, transaction db.DatabaseInstance) (bool, error) {
	result, err := service.Repository.Destroy(id, transaction)

	return result, err
}
