package interfaces

import (
	"events-system/infrastructure/providers/db"
	"events-system/internal/dto"
	entities "events-system/internal/entity"
	"time"
)

type UserService interface {
	Create(username string, transaction db.DatabaseInstance) (*entities.User, error)
	Find(options map[string]interface{}) (*[]entities.User, error)
	Update(id string, username string, transaction db.DatabaseInstance) (*entities.User, error)
	Delete(id string, transaction db.DatabaseInstance) (bool, error)
}

type EventService interface {
	CreateEvent(data dto.CreateEventDTO) (*dto.OutputEvent, error)
}

type AccountService interface {
	Create(data any, transaction db.DatabaseInstance) (*entities.Account, error)
	CheckAccount(accountId int64) (*entities.Account, error)
}

type TaskService interface {
	GetListOfTodayTasks() (*[]entities.Task, error)
	Create(
		data entities.CreateTaskData,
		transaction db.DatabaseInstance,
	) (*entities.Task, error)
	GenerateTimesForTasks(eventDate time.Time) []entities.TaskSliceEvent
}
