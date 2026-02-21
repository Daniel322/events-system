package vo

import "events-system/pkg/utils"

type EventType int

const (
	HappyBirthday EventType = iota
	Reminder
)

var SUPPORTED_EVENT_TYPES = map[EventType]string{
	HappyBirthday: "hb",
	Reminder:      "reminder",
}

func (event EventType) String() string {
	return SUPPORTED_EVENT_TYPES[event]
}

func NewEventType(s string) (EventType, error) {
	switch s {
	case "hb":
		return EventType(0), nil
	case "reminder":
		return EventType(1), nil
	default:
		return EventType(-1), utils.GenerateError("EventType", "invalid event type")
	}
}
