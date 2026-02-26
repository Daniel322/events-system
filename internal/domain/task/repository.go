package task

import (
	"context"
	"events-system/interfaces"
	"events-system/internal/components"
	"events-system/pkg/utils"

	"github.com/google/uuid"
)

type TaskRepo struct {
	components.Factory
}

var Repository *TaskRepo

func InitRepo(repo interfaces.Repository) {
	Repository = &TaskRepo{
		Factory: *components.NewFactory("Event", repo),
	}
}

func (r TaskRepo) Save(ctx context.Context, value interface{}) error {
	ctx = context.WithValue(ctx, "tableName", "tasks")

	return r.Repository.Save(ctx, value)
}

func (r TaskRepo) FindOne(ctx context.Context, options map[string]interface{}) (*Plain, error) {
	tasks := new([]Plain)
	ctx = context.WithValue(ctx, "tableName", "tasks")
	ctx = context.WithValue(ctx, "ptr", tasks)
	err := r.Repository.Find(ctx, options)

	if err != nil {
		return nil, utils.GenerateError("TaskRepo FindOne", err.Error())
	}
	if len(*tasks) == 0 {
		return nil, nil
	}

	task := (*tasks)[0]

	return &task, nil
}

func (r TaskRepo) Destroy(ctx context.Context, id string) error {
	ctx = context.WithValue(ctx, "tableName", "tasks")
	uuid, err := uuid.Parse(id)

	if err != nil {
		return utils.GenerateError("TaskRepo Destroy", err.Error())
	}

	return r.Repository.Destroy(ctx, interfaces.DestroyOptions{ID: uuid, Table: "tasks"})
}
