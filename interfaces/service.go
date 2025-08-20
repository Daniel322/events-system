package interfaces

import (
	"events-system/infrastructure/providers/db"
	"events-system/internal/dto"
	entities "events-system/internal/entity"
)

type IUserService interface {
	CreateUserWithAccount(data dto.UserDataDTO) (*dto.OutputUser, error)
	GetUser(id string) (*dto.OutputUser, error)
}

type IEventService interface {
	CreateEvent(data dto.CreateEventDTO) (*dto.OutputEvent, error)
}

type IAccountService interface {
	Create(data entities.Account, transaction db.DatabaseInstance) (*entities.Account, error)
	CheckAccount(accountId int64) (*entities.Account, error)
}
