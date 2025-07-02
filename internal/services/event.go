package services

import (
	"events-system/internal/domain"
	"events-system/internal/providers/db"
	"events-system/internal/repositories"
	"events-system/internal/utils"
	"strings"
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
	Info      string
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

	event, err := es.eventRepository.Create(domain.CreateEventData{
		UserId:       data.UserId,
		Info:         data.Info,
		Date:         data.Date,
		Providers:    strings.Split(string(data.Providers), ","),
		NotifyLevels: []string{"month", "week", "tomorrow", "today"},
	}, transaction)

	if err != nil {
		transaction.Rollback()
		return nil, utils.GenerateError(es.Name, err.Error())
	}

	timesForTask := make([]time.Time, 0)

	tasks := make([]domain.Task, 0)

	// TODO: fix dates, now make invalid dates for tasks
	// also need to right set of task type
	timesForTask = append(timesForTask, data.Date)
	timesForTask = append(timesForTask, data.Date.Add(-(time.Hour * 24)))
	timesForTask = append(timesForTask, data.Date.Add(-(time.Hour * 24 * 7)))
	timesForTask = append(timesForTask, data.Date.Add(-(time.Hour * 24 * 30)))

	for _, timeValue := range timesForTask {
		uuidV, _, err := utils.ParseId(data.AccountId)

		if err != nil {
			transaction.Rollback()
			return nil, utils.GenerateError(es.Name, err.Error())
		}

		task, err := es.taskRepository.Create(
			domain.CreateTaskData{
				EventId:   event.ID,
				AccountId: uuidV,
				Date:      timeValue,
				Type:      "today",
				Provider:  "telegram",
			},
			transaction,
		)

		if err != nil {
			transaction.Rollback()
			return nil, utils.GenerateError(es.Name, err.Error())
		}

		tasks = append(tasks, *task)
	}

	transaction.Commit()

	return &Event{
		ID:           event.ID,
		UserId:       event.UserId,
		Info:         event.Info,
		Date:         event.Date,
		NotifyLevels: event.NotifyLevels,
		Providers:    data.Providers,
		CreatedAt:    event.CreatedAt,
		UpdatedAt:    event.UpdatedAt,
		Tasks:        tasks,
	}, nil
}
