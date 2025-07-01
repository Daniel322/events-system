package services

import "events-system/internal/domain"

type IUserService interface {
	CreateUser(data CreateUserData) (*User, error)
	GetUser(id string) (*User, error)
}

type IEventService interface {
	CreateEvent(data CreateEventData) (*Event, error)
}

type IAccountService interface {
	CheckAccount(accountId int64) (*domain.Account, error)
}
