package interfaces

import (
	"context"

	"github.com/google/uuid"
)

type DestroyOptions struct {
	ID    uuid.UUID
	Table string
}

type Repository interface {
	Save(ctx context.Context, value interface{}) error
	Destroy(ctx context.Context, options DestroyOptions) error
	Find(
		ctx context.Context,
		options map[string]interface{},
	) error
	CreateTransaction(ctx context.Context) context.Context
	Commit(ctx context.Context)
	Rollback(ctx context.Context)
}
