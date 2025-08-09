package interfaces

import (
	"events-system/internal/domain"
	"events-system/internal/dto"
)

type IUserService interface {
	CreateUser(data dto.UserDataDTO) (*dto.OutputUser, error)
	GetUser(id string) (*dto.OutputUser, error)
}

type IEventService interface {
	CreateEvent(data dto.CreateEventDTO) (*dto.OutputEvent, error)
}

type IAccountService interface {
	CheckAccount(accountId int64) (*domain.Account, error)
}
