package queries

import (
	"context"
	"events-system/internal/domain/account"
	"events-system/internal/domain/event"
	"events-system/internal/domain/user"
	"events-system/pkg/utils"
	"log"
	"os"
)

type IGetUser struct {
	logger    *log.Logger
	userRepo  *user.UserRepo
	accRepo   *account.AccRepo
	eventRepo *event.EventsRepo
}

var GetUser *IGetUser

func InitGetUser() {
	var logger = log.New(os.Stdout, "GetUser ", log.LstdFlags)

	GetUser = &IGetUser{
		userRepo:  user.Repository,
		accRepo:   account.Repository,
		eventRepo: event.Repository,
		logger:    logger,
	}
}

func (this IGetUser) Run(ctx context.Context, id string) (*user.Plain, error) {
	findOptions := make(map[string]interface{})
	findOptions["id"] = id

	user, err := this.userRepo.FindOne(ctx, findOptions)

	if err != nil {
		return nil, utils.GenerateError("GetUser", err.Error())
	}

	// accOptions := make(map[string]interface{})
	// accOptions["user_id"] = id

	// accs, err := this.AccRepo.Repository.Find(ctx, accOptions)
	// if err != nil {
	// 	return nil, utils.GenerateError("GetUser", err.Error())
	// }

	// for _, v := range *accs {
	// 	user.AddAccount(*(v.(*account.Entity)))
	// }

	return user, nil
}
