package commands

import (
	"context"
	"events-system/internal/domain/account"
	"events-system/internal/domain/event"
	"events-system/internal/domain/task"
	"events-system/internal/domain/user"
	"events-system/pkg/utils"
	"events-system/pkg/vo"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type CreateEvent struct {
	Logger    *log.Logger
	UserRepo  *user.UserRepo
	AccRepo   *account.AccRepo
	EventRepo *event.EventsRepo
	TaskRepo  *task.TaskRepo
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

func NewCreateEvent(
	userRepo *user.UserRepo,
	accRepo *account.AccRepo,
	eventRepo *event.EventsRepo,
	taskRepo *task.TaskRepo,
) *CreateEvent {
	var logger = log.New(os.Stdout, "CreateEvent ", log.LstdFlags)

	return &CreateEvent{
		UserRepo:  userRepo,
		AccRepo:   accRepo,
		EventRepo: eventRepo,
		TaskRepo:  taskRepo,
		Logger:    logger,
	}
}

func (this CreateEvent) Validate(data CreateEventData) (*CreateEventState, error) {
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

	accId, err := uuid.Parse(data.UserId)

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

func (this CreateEvent) Run(
	ctx context.Context,
	state *CreateEventState,
) (*event.Entity, error) {
	// TODO: add transaction
	event := event.New(
		state.Info,
		state.Date,
		state.EventType,
		state.NotifyLevels,
		state.Providers,
		state.UserId,
	)

	err := this.EventRepo.Save(ctx, event.ToPlain())

	if err != nil {
		return nil, utils.GenerateError("Create event", err.Error())
	}

	// TODO: add generate tasks info and create tasks

	return &event, nil
}
