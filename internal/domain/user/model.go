package user

import (
	"events-system/pkg/vo"
)

type Model struct {
	username vo.NonEmptyString
}

func (m Model) Username() string {
	return m.username.Val()
}

func newModel(username vo.NonEmptyString) *Model {
	return &Model{
		username: username,
	}
}
