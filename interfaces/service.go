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

type AccountService interface {
	Create(data dto.CreateAccountData, transaction db.DatabaseInstance) (*entities.User, error)
	Find(options map[string]interface{}) (*[]entities.User, error)
	Update(id string, data dto.UpdateAccountData, transaction db.DatabaseInstance) (*entities.User, error)
	Delete(id string, transaction db.DatabaseInstance) (bool, error)
}

type EventService interface {
	Create(data dto.CreateEventData, transaction db.DatabaseInstance) (*entities.Event, error)
	Find(options map[string]interface{}) (*[]entities.Event, error)
	Update(id string, data dto.UpdateEventData, transaction db.DatabaseInstance) (*entities.Event, error)
	Delete(id string, transaction db.DatabaseInstance) (bool, error)
}

type TaskService interface {
	Find(options map[string]interface{}) (*[]entities.Task, error)
	Delete(id string, transaction db.DatabaseInstance) (bool, error)
	Create(data dto.CreateTaskData, transaction db.DatabaseInstance) (*entities.Event, error)
	Update(id string, date time.Time, transaction db.DatabaseInstance) (*entities.Event, error)
}
