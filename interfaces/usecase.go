package interfaces

import (
	"events-system/internal/dto"
	entities "events-system/internal/entity"
)

type InternalUsecase interface {
	CreateUser(data dto.CreateUserInput) (*dto.OutputUser, error)
	GetUser(id string) (*dto.OutputUser, error)
	CheckTGAccount(accountId int64) (*entities.Account, error)
	GetListOfTodayTasks() (*[]entities.Task, error)
	CreateEvent(data dto.CreateEventDTO) (*dto.OutputEvent, error)
	ExecTask(taskId string) (*dto.InfoAboutTaskForTgProvider, error)
}
