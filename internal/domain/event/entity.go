package event

import (
	"events-system/interfaces"
	"events-system/internal/components/vo"
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	interfaces.Entity
	Model
	UserId uuid.UUID
}

type Plain struct {
	ID           string
	UserId       string
	Info         string
	Date         time.Time
	Type         string
	NotifyLevels []byte
	Providers    []byte
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func New(
	info vo.NonEmptyString,
	date time.Time,
	eventType vo.EventType,
	notifyLevels vo.JsonField,
	providers vo.JsonField,
	userId uuid.UUID,
) Entity {
	return Entity{
		interfaces.NewEntity(),
		newModel(
			info,
			date,
			eventType,
			notifyLevels,
			providers,
		),
		userId,
	}
}

func (e Entity) ToPlain() Plain {
	notify, _ := e.notifyLevels.Value()
	providers, _ := e.providers.Value()
	return Plain{
		ID:           e.ID.String(),
		UserId:       e.UserId.String(),
		Info:         e.info.Val(),
		Date:         e.date,
		NotifyLevels: notify.([]byte),
		Providers:    providers.([]byte),
		Type:         e.eventType.String(),
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
	}
}
