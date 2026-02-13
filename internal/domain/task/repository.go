package task

import (
	"context"
	"events-system/interfaces"
	"events-system/internal/components"
)

type TaskRepo struct {
	components.Factory
}

func NewTaskRepo(repo interfaces.RepositoryV2) *TaskRepo {
	return &TaskRepo{
		Factory: *components.NewFactory("Event", repo),
	}
}

func (r TaskRepo) Save(ctx context.Context, value interface{}) error {
	ctx = context.WithValue(ctx, "tableName", "tasks")

	return r.Repository.Save(ctx, value)
}
