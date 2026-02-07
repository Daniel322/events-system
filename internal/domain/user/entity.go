package user

import (
	"events-system/interfaces"
	"events-system/pkg/vo"
	"time"
)

type Entity struct {
	interfaces.Entity
	*Model
}

type Plain struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
}

func New(username vo.NonEmptyString) Entity {
	return Entity{
		interfaces.NewEntity(),
		newModel(username),
	}
}

func (e Entity) ToPlain() Plain {
	return Plain{
		ID:        e.ID.String(),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		Username:  e.Username(),
	}
}
