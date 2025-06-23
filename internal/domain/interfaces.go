package domain

type IAccountFactory interface {
	CreateAccount(data CreateAccountData) (*Account, error)
	UpdateAccount(acc *Account, data UpdateAccountData) (*Account, error)
}

type IUserFactory interface {
	CreateUser(data UserData) (*User, error)
	UpdateUser(user *User, data UserData) (*User, error)
}
