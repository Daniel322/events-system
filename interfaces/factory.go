package interfaces

import entities "events-system/internal/entity"

type UserFactory interface {
	Create(username string) (*entities.User, error)
	Update(user *entities.User, username string) (*entities.User, error)
}

type AccountFactory interface {
	Create(data entities.CreateAccountData) (*entities.Account, error)
	Update(acc *entities.Account, data entities.UpdateAccountData) (*entities.Account, error)
}

type EventFactory interface {
	Create(data entities.CreateEventData) (*entities.Event, error)
	Update(event *entities.Event, data entities.UpdateEventData) (*entities.Event, error)
}

type TaskFactory interface {
	Create(data entities.CreateTaskData) (*entities.Task, error)
	Update(task *entities.Task, data entities.UpdateTaskData) (*entities.Task, error)
}
