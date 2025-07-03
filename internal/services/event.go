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

type TaskSliceEvent struct {
	Date time.Time
	Type string
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

	timesForTask := make([]TaskSliceEvent, 0)

	tasks := make([]domain.Task, 0)

	today := time.Now()
	todayYear := today.Year()
	eventDateYear := data.Date.Year()
	currentEventInThatYear := data.Date
	if eventDateYear < todayYear {
		currentEventInThatYear = time.Date(
			todayYear,
			data.Date.Month(),
			data.Date.Day(),
			data.Date.Hour(),
			data.Date.Minute(),
			data.Date.Second(),
			data.Date.Nanosecond(),
			data.Date.Location(),
		)
		// if event in that year before today
		if currentEventInThatYear.Compare(today) == -1 {
			currentEventInThatYear = time.Date(
				todayYear+1,
				data.Date.Month(),
				data.Date.Day(),
				data.Date.Hour(),
				data.Date.Minute(),
				data.Date.Second(),
				data.Date.Nanosecond(),
				data.Date.Location(),
			)
		}
	}
	timesForTask = append(timesForTask, TaskSliceEvent{Date: currentEventInThatYear, Type: "today"})
	timesForTask = append(timesForTask, TaskSliceEvent{Date: currentEventInThatYear.Add(-(time.Hour * 24)), Type: "tomorrow"})
	timesForTask = append(timesForTask, TaskSliceEvent{Date: currentEventInThatYear.Add(-(time.Hour * 24 * 7)), Type: "week"})
	timesForTask = append(timesForTask, TaskSliceEvent{Date: currentEventInThatYear.Add(-(time.Hour * 24 * 30)), Type: "month"})

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
				Date:      timeValue.Date,
				Type:      timeValue.Type,
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
