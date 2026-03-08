package commands

import (
	"context"
	pg_db "events-system/infrastructure/db/adapters/postgres"
	"events-system/internal/domain/account"
	"events-system/internal/domain/event"
	"events-system/internal/domain/task"
	"events-system/pkg/utils"
	"log"

	"github.com/google/uuid"
)

type IExectask struct {
	logger    *log.Logger
	taskRepo  *task.TaskRepo
	eventRepo *event.EventsRepo
	accRepo   *account.AccRepo
}

type ExecTaskData struct {
	Id string
}

type ExecTaskState struct {
	id string
}

type ExecTaskResult struct {
	ChatId string
	Text   string
}

var ExecTask *IExectask

func InitExecTask() {
	var logger = log.New(log.Writer(), "ExecTask ", log.LstdFlags)

	ExecTask = &IExectask{
		taskRepo:  task.Repository,
		eventRepo: event.Repository,
		accRepo:   account.Repository,
		logger:    logger,
	}
}

func (this IExectask) Validate(data ExecTaskData) (*ExecTaskState, error) {
	state := ExecTaskState{}

	state.id = data.Id

	return &state, nil
}

func (this IExectask) Run(ctx context.Context, state *ExecTaskState) (*ExecTaskResult, error) {
	if ctx.Value("transaction") == nil {
		transaction := pg_db.Adapter.CreateTransaction()

		ctx = context.WithValue(ctx, "transaction", transaction)
	}
	// find task by id from state
	currentTask, err := this.taskRepo.FindOne(ctx, map[string]interface{}{"id": state.id})

	if err != nil {
		return nil, utils.GenerateError("ExecTask.Run", err.Error())
	}

	currentEvent, err := this.eventRepo.FindOne(ctx, map[string]interface{}{"id": currentTask.EventId})

	if err != nil {
		return nil, utils.GenerateError("ExecTask.Run", err.Error())
	}

	currentAcc, err := this.accRepo.FindOne(ctx, map[string]interface{}{"id": currentTask.AccountId})

	if err != nil {
		return nil, utils.GenerateError("ExecTask.Run", err.Error())
	}

	// delete task
	this.taskRepo.Destroy(ctx, state.id)
	// create new task with same data but date = date + time of notify level (today, tomorrow, etc.)
	taskType, _ := task.NewTaskType(currentTask.Type)
	provider, _ := task.NewTaskProvider(currentTask.Provider)
	accId, _ := uuid.Parse(currentTask.AccountId)
	eventId, _ := uuid.Parse(currentTask.EventId)
	newTask := task.New(currentTask.Date.AddDate(1, 0, 0), taskType, provider, accId, eventId)

	err = this.taskRepo.Save(ctx, newTask.ToPlain())

	if err != nil {
		return nil, utils.GenerateError("ExecTask.Run", err.Error())
	}

	text := "Attention!" + " For " + (*currentTask).Type + " in " + (*currentEvent).Date.Format("2 January") + " will be event " + (*currentEvent).Info

	return &ExecTaskResult{ChatId: currentAcc.AccountId, Text: text}, nil
}
