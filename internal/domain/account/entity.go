package account

import (
	"events-system/interfaces"
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	interfaces.Entity
	Model
	UserId uuid.UUID
}

type Plain struct {
	ID        string
	UserId    string
	AccountId string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(value AccountValue, acctype AccountType, userId uuid.UUID) Entity {
	return Entity{
		interfaces.NewEntity(),
		newModel(value, acctype),
		userId,
	}
}

func (e Entity) ToPlain() Plain {
	return Plain{
		ID:        e.ID.String(),
		UserId:    e.UserId.String(),
		AccountId: e.value.Val(),
		Type:      e.acctype.String(),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
