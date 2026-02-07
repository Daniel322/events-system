package user

import (
	"events-system/internal/domain/account"
	"events-system/internal/domain/event"
	"events-system/pkg/vo"
)

type Model struct {
	username vo.NonEmptyString
	accounts *[]account.Model
	events   *[]event.Model
}

func (m Model) Username() string {
	return m.username.Val()
}

func newModel(username vo.NonEmptyString) *Model {
	return &Model{
		username: username,
		accounts: new([]account.Model),
		events:   new([]event.Model),
	}
}

func (m *Model) AddAccount(acc account.Model) {
	*m.accounts = append(*m.accounts, acc)
}

func (m Model) Accounts() []account.Model {
	return *m.accounts
}
