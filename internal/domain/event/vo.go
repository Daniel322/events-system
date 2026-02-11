package event

import (
	"events-system/pkg/utils"
	"events-system/pkg/vo"
	"slices"
)

// EVENT TYPE

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

// Notify levels

type NotifyLevels vo.JsonField

var SUPPORTED_NOTIFY_LEVELS = []string{"today", "tomorrow", "month", "week"}

func NewNotifyLevels(values []string) (vo.JsonField, error) {
	result := vo.JsonField{}
	for _, v := range values {
		if ok := slices.Contains(SUPPORTED_NOTIFY_LEVELS, v); !ok {
			return vo.JsonField{}, utils.GenerateError("NotifyLevels", "invalid notify level")
		} else {
			result = append(result, v)
		}
	}

	return result, nil
}

// Providers

type Providers vo.JsonField

var SUPPORTED_PROVIDERS = []string{"mail", "telegram"}

func NewProviders(values []string) (vo.JsonField, error) {
	result := vo.JsonField{}
	for _, v := range values {
		if ok := slices.Contains(SUPPORTED_PROVIDERS, v); !ok {
			return vo.JsonField{}, utils.GenerateError("Providers", "invalid provider")
		} else {
			result = append(result, v)
		}
	}

	return result, nil
}
