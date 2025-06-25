package domain

type Factory[Entity any, CreateData any, UpdateData any] interface {
	Create(data CreateData) (*Entity, error)
	Update(t *Entity, data UpdateData) (*Entity, error)
}

type IUserFactory interface {
	Factory[User, UserData, UserData]
}

type IAccountFactory interface {
	Factory[Account, CreateAccountData, UpdateAccountData]
}
