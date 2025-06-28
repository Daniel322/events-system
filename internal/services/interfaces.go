package services

type IUserService interface {
	CreateUser(data CreateUserData) (*User, error)
	GetUser(id string) (*User, error)
}
