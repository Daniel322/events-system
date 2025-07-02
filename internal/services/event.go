package services

import (
	"events-system/internal/domain"
	"events-system/internal/providers/db"
	"events-system/internal/repositories"
	"events-system/internal/utils"
	"log"
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

	log.Println(es.Name, "DATA:", data)

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

	// log.Println(data.Date, data.Date.Add(-(time.Hour * 24)), data.Date.Add(-(time.Hour * 24 * 7)))
	log.Println(timesForTask)
	timesForTask = append(timesForTask, data.Date)
	log.Println(timesForTask)
	timesForTask = append(timesForTask, data.Date.Add(-(time.Hour * 24)))
	log.Println(timesForTask)
	timesForTask = append(timesForTask, data.Date.Add(-(time.Hour * 24 * 7)))
	log.Println(timesForTask)
	timesForTask = append(timesForTask, data.Date.Add(-(time.Hour * 24 * 30)))
	log.Println(timesForTask)

	for _, timeValue := range timesForTask {
		uuidV, _, err := utils.ParseId(data.AccountId)

		log.Println(es.Name, "uuidV:", uuidV)

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

		log.Println(es.Name, "task:", task)

		if err != nil {
			transaction.Rollback()
			return nil, utils.GenerateError(es.Name, err.Error())
		}

		tasks = append(tasks, *task)
	}

	log.Println(es.Name, "tasks after loop:", tasks)

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
