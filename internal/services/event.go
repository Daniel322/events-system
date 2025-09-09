package services

import (
	"events-system/infrastructure/providers/db"
	"events-system/interfaces"
	"events-system/internal/dto"
	entities "events-system/internal/entity"
	"events-system/pkg/utils"
	"time"

	"github.com/google/uuid"
)

type EventService struct {
	Name       string
	Repository interfaces.Repository[entities.Event]
}

const (
	INVALID_INFO          = "invalid info"
	INVALID_DATE          = "invalid date"
	INVALID_NOTIFY_LEVELS = "error on parse notify levels"
	INVALID_PROVIDERS     = "error on parse providers"
)

func NewEventService(repository interfaces.Repository[entities.Event]) *EventService {
	return &EventService{
		Name:       "EventService",
		Repository: repository,
	}
}

func (service *EventService) checkInfo(value string) error {
	if len(value) == 0 {
		return utils.GenerateError(service.Name, INVALID_INFO)
	}

	return nil
}

func (service *EventService) checkDate(value time.Time) error {
	if value.IsZero() {
		return utils.GenerateError(service.Name, INVALID_DATE)
	}
	return nil
}

// TODO: use DRY and optimize code in that and next funcs
func (service *EventService) checkProviders(value entities.Providers) error {
	var dest interface{}
	err := value.Scan(dest)

	if err != nil {
		return err
	}

	_, err = value.Value()

	if err != nil {
		return err
	}

	return nil
}

func (service *EventService) checkNotifyLevels(value entities.NotifyLevel) (entities.NotifyLevel, error) {
	if len(value) == 0 {
		return entities.NOTIFY_LEVELS, nil
	} else {
		var dest interface{}
		err := value.Scan(dest)

		if err != nil {
			return nil, err
		}

		_, err = value.Value()

		if err != nil {
			return nil, err
		}
	}
	return value, nil
}

func (service *EventService) Find(options map[string]interface{}) (*[]entities.Event, error) {
	results, err := service.Repository.Find(options)

	return results, err
}

func (service *EventService) FindOne(options map[string]interface{}) (*entities.Event, error) {
	events, err := service.Find(options)

	if err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	if len(*events) == 0 {
		return nil, utils.GenerateError(service.Name, "current account not found")
	}

	return &(*events)[0], nil
}

func (service *EventService) Delete(id string, transaction db.DatabaseInstance) (bool, error) {
	result, err := service.Repository.Destroy(id, transaction)

	return result, err
}

func (service *EventService) Create(data dto.CreateEventData, transaction db.DatabaseInstance) (*entities.Event, error) {
	var id uuid.UUID = uuid.New()

	if err := uuid.Validate(data.UserId.String()); err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	err := service.checkInfo(data.Info)

	if err != nil {
		return nil, err
	}

	err = service.checkDate(data.Date)

	if err != nil {
		return nil, err
	}

	notifyLevels, err := service.checkNotifyLevels(data.NotifyLevels)

	if err != nil {
		return nil, err
	}

	err = service.checkProviders(data.Providers)

	if err != nil {
		return nil, err
	}

	event := &entities.Event{
		ID:           id,
		UserId:       data.UserId,
		Info:         data.Info,
		Date:         data.Date,
		NotifyLevels: notifyLevels,
		Providers:    data.Providers,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	event, err = service.Repository.Save(*event, transaction)

	return event, err
}

func (service *EventService) Update(
	id string,
	data dto.UpdateEventData,
	transaction db.DatabaseInstance,
) (*entities.Event, error) {
	findOptions := make(map[string]interface{})
	findOptions["id"] = id

	currentEvent, err := service.FindOne(findOptions)

	if err != nil {
		return nil, utils.GenerateError(service.Name, err.Error())
	}

	if isInvalidInfo := service.checkInfo(data.Info); isInvalidInfo == nil {
		currentEvent.Info = data.Info
	}
	if isInvalidDate := service.checkDate(data.Date); isInvalidDate == nil {
		currentEvent.Date = data.Date
	}
	if notifyLevels, isInvalidNotifyLevels := service.checkNotifyLevels(data.NotifyLevels); isInvalidNotifyLevels == nil {
		currentEvent.NotifyLevels = notifyLevels
	}
	if isInvalidProviders := service.checkProviders(data.Providers); isInvalidProviders == nil {
		currentEvent.Providers = data.Providers
	}

	currentEvent.UpdatedAt = time.Now()

	currentEvent, err = service.Repository.Save(*currentEvent, transaction)

	return currentEvent, err
}
