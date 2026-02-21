package commands

import (
	"context"
	"events-system/internal/domain/account"
	"events-system/internal/domain/event"
	"events-system/internal/domain/task"
	"events-system/pkg/utils"
	"log"

	"github.com/google/uuid"
)

type Exectask struct {
	Logger    *log.Logger
	TaskRepo  *task.TaskRepo
	EventRepo *event.EventsRepo
	AccRepo   *account.AccRepo
}

type ExecTaskData struct {
	id string
}

type ExecTaskState struct {
	id string
}

type ExecTaskResult struct {
	ChatId string
	Text   string
}

func NewExecTask(
	taskRepo *task.TaskRepo,
	eventRepo *event.EventsRepo,
	accRepo *account.AccRepo,
) *Exectask {
	var logger = log.New(log.Writer(), "ExecTask ", log.LstdFlags)

	return &Exectask{
		TaskRepo:  taskRepo,
		EventRepo: eventRepo,
		AccRepo:   accRepo,
		Logger:    logger,
	}
}

func (this Exectask) Validate(data ExecTaskData) (*ExecTaskState, error) {
	state := ExecTaskState{}

	state.id = data.id

	return &state, nil
}

func (this Exectask) Run(ctx context.Context, state *ExecTaskState) (*ExecTaskResult, error) {
	// find task by id from state
	currentTask, err := this.TaskRepo.FindOne(ctx, map[string]interface{}{"id": state.id})

	if err != nil {
		return nil, utils.GenerateError("ExecTask.Run", err.Error())
	}

	currentEvent, err := this.EventRepo.FindOne(ctx, map[string]interface{}{"id": currentTask.EventId})

	if err != nil {
		return nil, utils.GenerateError("ExecTask.Run", err.Error())
	}

	currentAcc, err := this.AccRepo.FindOne(ctx, map[string]interface{}{"id": currentTask.AccountId})

	if err != nil {
		return nil, utils.GenerateError("ExecTask.Run", err.Error())
	}

	// delete task
	this.TaskRepo.Destroy(ctx, state.id)
	// create new task with same data but date = date + time of notify level (today, tomorrow, etc.)
	taskType, _ := task.NewTaskType(currentTask.Type)
	provider, _ := task.NewTaskProvider(currentTask.Provider)
	accId, _ := uuid.Parse(currentTask.AccountId)
	eventId, _ := uuid.Parse(currentTask.EventId)
	newTask := task.New(currentTask.Date.AddDate(1, 0, 0), taskType, provider, accId, eventId)

	err = this.TaskRepo.Save(ctx, newTask)

	if err != nil {
		return nil, utils.GenerateError("ExecTask.Run", err.Error())
	}

	text := "Attention!" + " For " + currentTask.Type + " in " + currentEvent.Date.Format("01-02") + " will be event " + currentEvent.Info

	return &ExecTaskResult{ChatId: currentAcc.AccountId, Text: text}, nil
}
