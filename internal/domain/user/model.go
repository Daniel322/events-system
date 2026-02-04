package user

import (
	"events-system/internal/domain/account"
	"events-system/internal/domain/event"
	"events-system/pkg/vo"
)

type Model struct {
	username vo.NonEmptyString
	accounts []account.Model
	events   []event.Model
}

func (m Model) Username() string {
	return m.username.Val()
}

func newModel(username vo.NonEmptyString) Model {
	return Model{
		username: username,
	}
}
