package event

import (
	"events-system/pkg/vo"
	"time"
)

type Model struct {
	info         vo.NonEmptyString
	date         time.Time
	eventType    EventType
	notifyLevels vo.JsonField
	providers    vo.JsonField
}

func (m Model) Type() string {
	return m.eventType.String()
}

func newModel(
	info vo.NonEmptyString,
	date time.Time,
	eventType EventType,
	notifyLevels vo.JsonField,
	providers vo.JsonField,
) Model {
	return Model{
		info:         info,
		date:         date,
		eventType:    eventType,
		notifyLevels: notifyLevels,
		providers:    providers,
	}
}
