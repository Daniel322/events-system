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

type GetUser struct {
	Logger    *log.Logger
	UserRepo  *user.UserRepo
	AccRepo   *account.AccRepo
	EventRepo *event.EventsRepo
}

func NewGetUser(
	userRepo *user.UserRepo,
	accRepo *account.AccRepo,
	eventRepo *event.EventsRepo,
) *GetUser {
	var logger = log.New(os.Stdout, "GetUser ", log.LstdFlags)

	return &GetUser{
		UserRepo:  userRepo,
		AccRepo:   accRepo,
		EventRepo: eventRepo,
		Logger:    logger,
	}
}

func (this GetUser) Run(ctx context.Context, id string) (*user.Entity, error) {
	findOptions := make(map[string]interface{})
	findOptions["id"] = id

	user, err := this.UserRepo.FindOne(ctx, findOptions)

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
