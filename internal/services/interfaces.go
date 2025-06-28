package services

type IUserService interface {
	CreateUser(data CreateUserData) (*User, error)
	GetUser(id string) (*User, error)
}

type IEventService interface {
	CreateEvent(data CreateEventData) (*Event, error)
}
