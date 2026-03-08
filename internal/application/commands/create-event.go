package commands

import (
	"context"
	pg_db "events-system/infrastructure/db/adapters/postgres"
	"events-system/internal/components/vo"
	"events-system/internal/domain/account"
	"events-system/internal/domain/event"
	"events-system/internal/domain/task"
	"events-system/internal/domain/user"
	"events-system/pkg/utils"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type ICreateEvent struct {
	logger    *log.Logger
	userRepo  *user.UserRepo
	accRepo   *account.AccRepo
	eventRepo *event.EventsRepo
	taskRepo  *task.TaskRepo
}

type CreateEventData struct {
	UserId       string
	AccId        string
	Info         string
	Date         time.Time
	NotifyLevels []string
	Providers    []string
}

type CreateEventState struct {
	UserId       uuid.UUID
	AccId        uuid.UUID
	Info         vo.NonEmptyString
	Date         time.Time
	NotifyLevels vo.JsonField
	Providers    vo.JsonField
	EventType    vo.EventType
}

var CreateEvent *ICreateEvent

func InitCreateEvent() {
	var logger = log.New(os.Stdout, "CreateEvent ", log.LstdFlags)

	CreateEvent = &ICreateEvent{
		userRepo:  user.Repository,
		accRepo:   account.Repository,
		eventRepo: event.Repository,
		taskRepo:  task.Repository,
		logger:    logger,
	}
}

func (this ICreateEvent) Validate(data CreateEventData) (*CreateEventState, error) {
	state := CreateEventState{}

	state.Date = data.Date

	info, err := vo.NewNonEmptyString(data.Info)

	if err != nil {
		return nil, utils.GenerateError("CreateEvent.Validate", err.Error())
	}

	state.Info = info

	userId, err := uuid.Parse(data.UserId)

	if err != nil {
		return nil, utils.GenerateError("CreateEvent.Validate", err.Error())
	}

	state.UserId = userId

	accId, err := uuid.Parse(data.AccId)

	if err != nil {
		return nil, utils.GenerateError("CreateEvent.Validate", err.Error())
	}

	state.AccId = accId

	not, err := event.NewNotifyLevels(data.NotifyLevels)

	if err != nil {
		return nil, utils.GenerateError("CreateEvent.Validate", err.Error())
	}

	state.NotifyLevels = not

	providers, err := event.NewProviders(data.Providers)

	if err != nil {
		return nil, utils.GenerateError("CreateEvent.Validate", err.Error())
	}

	state.Providers = providers

	state.EventType, _ = vo.NewEventType("hb")

	return &state, nil
}

func (this ICreateEvent) Run(
	ctx context.Context,
	state *CreateEventState,
) (*event.Entity, error) {

	if ctx.Value("transaction") == nil {
		transaction := pg_db.Adapter.CreateTransaction()

		ctx = context.WithValue(ctx, "transaction", transaction)
	}

	event := event.New(
		state.Info,
		state.Date,
		state.EventType,
		state.NotifyLevels,
		state.Providers,
		state.UserId,
	)

	err := this.eventRepo.Save(ctx, event.ToPlain())

	if err != nil {
		return nil, utils.GenerateError("Create event", err.Error())
	}

	var TASKS_TYPES = map[string]time.Duration{
		"today":    0,
		"tomorrow": time.Hour * 24,
		"week":     time.Hour * 168,
		"month":    time.Hour * 720,
	}

	// tasks := make(
	// 	[]task.Entity,
	// 	0,
	// 	len(state.NotifyLevels)*len(state.Providers),
	// )
	for _, provider := range state.Providers {
		for _, level := range state.NotifyLevels {
			today := time.Now()
			todayYear := today.Year()
			currentEventInThatYear := time.Date(
				todayYear,
				state.Date.Month(),
				state.Date.Day(),
				state.Date.Hour(),
				state.Date.Minute(),
				state.Date.Second(),
				state.Date.Nanosecond(),
				state.Date.Location(),
			).Add(-TASKS_TYPES[level])

			if currentEventInThatYear.Compare(today) == -1 {
				currentEventInThatYear = currentEventInThatYear.AddDate(1, 0, 0)
			}

			taskType, _ := task.NewTaskType(level)

			taskProvider, _ := task.NewTaskProvider(provider)

			task := task.New(
				currentEventInThatYear,
				taskType,
				taskProvider,
				state.AccId,
				event.ID,
			)

			err = this.taskRepo.Save(ctx, task.ToPlain())

			// tasks = append(tasks, task)
		}
	}

	return &event, nil
}
