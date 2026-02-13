package task

import (
	"events-system/interfaces"
	"events-system/pkg/vo"
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	interfaces.Entity
	Model
	EventId   uuid.UUID
	AccountId uuid.UUID
}

type Plain struct {
	ID        string
	EventId   string
	AccountId string
	Provider  string
	Type      string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(
	date time.Time,
	taskType vo.EventType,
	provider TaskProvider,
	accountId uuid.UUID,
	eventId uuid.UUID,
) Entity {
	return Entity{
		interfaces.NewEntity(),
		newModel(
			date,
			taskType,
			provider,
		),
		eventId,
		accountId,
	}
}

func (e Entity) ToPlain() Plain {
	return Plain{
		ID:        e.ID.String(),
		EventId:   e.EventId.String(),
		AccountId: e.AccountId.String(),
		Date:      e.date,
		Type:      e.taskType.String(),
		Provider:  e.provider.String(),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
