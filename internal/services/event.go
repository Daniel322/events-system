package services

import (
	"events-system/internal/domain"
	"events-system/internal/providers/db"
	"events-system/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type EventService struct {
	Name            string
	DB              *db.Database
	eventRepository *repositories.Repository[domain.Event, domain.CreateEventData, domain.UpdateEventData]
	taskRepository  *repositories.Repository[domain.Task, domain.CreateTaskData, domain.UpdateTaskData]
}

type CreateEventData struct {
	AccountId string
	UserId    string
	Name      string
	Date      time.Time
	Providers []byte
}

type Event struct {
	ID           uuid.UUID
	UserId       uuid.UUID
	Info         string
	Date         time.Time
	NotifyLevels []byte
	Providers    []byte
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Tasks        []domain.Task
}

func NewEventService(
	db *db.Database,
	eventRepository *repositories.Repository[domain.Event, domain.CreateEventData, domain.UpdateEventData],
	taskRepository *repositories.Repository[domain.Task, domain.CreateTaskData, domain.UpdateTaskData],
) *EventService {
	return &EventService{
		Name:            "EventService",
		DB:              db,
		eventRepository: eventRepository,
		taskRepository:  taskRepository,
	}
}

func (es *EventService) CreateEvent(data CreateEventData) (*Event, error) {
	transaction := es.DB.CreateTransaction()

	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()

	// event, err := es.eventRepository.Create()

	return &Event{}, nil
}
