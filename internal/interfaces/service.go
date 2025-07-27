package interfaces

import (
	"events-system/internal/domain"
	"events-system/internal/structs"
)

type IUserService interface {
	CreateUser(data structs.CreateUserData) (*structs.User, error)
	GetUser(id string) (*structs.User, error)
}

type IEventService interface {
	CreateEvent(data structs.CreateEventData) (*structs.Event, error)
}

type IAccountService interface {
	CheckAccount(accountId int64) (*domain.Account, error)
}
