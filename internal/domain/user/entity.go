package user

import (
	"events-system/interfaces"
	"events-system/pkg/vo"
)

type Entity struct {
	interfaces.Entity
	Model
}

func New(username vo.NonEmptyString) Entity {
	return Entity{
		interfaces.NewEntity(),
		newModel(username),
	}
}
