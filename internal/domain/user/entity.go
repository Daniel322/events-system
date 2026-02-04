package user

import (
	"events-system/pkg/vo"
)

type Entity struct {
	vo.Entity
	Model
}

func New(username vo.NonEmptyString) Entity {
	return Entity{
		vo.NewEntity(),
		newModel(username),
	}
}
