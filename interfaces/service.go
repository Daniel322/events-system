package interfaces

import (
	"events-system/infrastructure/providers/db"
	"events-system/internal/dto"
	entities "events-system/internal/entity"
)

type UserService interface {
	CreateUserWithAccount(data dto.UserDataDTO) (*dto.OutputUser, error)
	GetUser(id string) (*dto.OutputUser, error)
}

type EventService interface {
	CreateEvent(data dto.CreateEventDTO) (*dto.OutputEvent, error)
}

type AccountService interface {
	Create(data entities.Account, transaction db.DatabaseInstance) (*entities.Account, error)
	CheckAccount(accountId int64) (*entities.Account, error)
}

type TaskService interface {
	GetListOfTodayTasks() (*[]entities.Task, error)
	Create(
		data entities.CreateTaskData,
		transaction db.DatabaseInstance,
	) (*entities.Task, error)
}
