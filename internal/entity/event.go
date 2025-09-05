package entities

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID     uuid.UUID
	UserId uuid.UUID
	Info   string
	Date   time.Time
	// TODO: think how move that annotation to other struct for keep clean domain entity
	NotifyLevels NotifyLevel `gorm:"type:jsonb"`
	Providers    Providers   `gorm:"type:jsonb"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

var NOTIFY_LEVELS = NotifyLevel{"month", "week", "tomorrow", "today"}

// func (ef *EventFactory) Update(event *Event, data UpdateEventData) (*Event, error) {
// 	dataValue := reflect.ValueOf(&data).Elem()
// 	var fields = 0

// 	if infoField := dataValue.FieldByName("Info"); infoField.IsValid() {
// 		event.Info = data.Info
// 		fields++
// 	}

// 	// TODO: fix after repair create event use case
// 	// if notifyField := dataValue.FieldByName("NotifyLevels"); notifyField.IsValid() {
// 	// 	event.NotifyLevels = data.NotifyLevels
// 	// 	fields++
// 	// }

// 	// if providersField := dataValue.FieldByName("Providers"); providersField.IsValid() {
// 	// 	event.Providers = data.Providers
// 	// 	fields++
// 	// }

// 	if dateField := dataValue.FieldByName("Date"); dateField.IsValid() {
// 		event.Date = data.Date
// 		fields++
// 	}

// 	if fields > 0 {
// 		event.UpdatedAt = time.Now()
// 	}

// 	return event, nil
// }
