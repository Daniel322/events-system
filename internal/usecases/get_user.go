package usecases

import (
	"events-system/internal/dto"
	"events-system/pkg/utils"
)

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
