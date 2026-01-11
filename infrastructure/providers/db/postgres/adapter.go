package pg_db

import (
	"context"
	"events-system/interfaces"
	"log"
	"os"

	"gorm.io/gorm"
)

type DbAdapter struct {
	Instance *gorm.DB
	Logger   *log.Logger
}

func NewDbAdapter(instance *gorm.DB) *DbAdapter {
	var logger = log.New(os.Stdout, "DB Adapter ", log.LstdFlags)

	return &DbAdapter{
		Instance: instance,
		Logger:   logger,
	}
}

func (adapter *DbAdapter) instance(ctx context.Context) *gorm.DB {
	if ctx.Value("transaction") != nil {
		return ctx.Value("transaction").(*gorm.DB)
	}

	return adapter.Instance
}

func (adapter *DbAdapter) Save(ctx context.Context, value interface{}) error {
	adapterContextForExecQuery := adapter.instance(ctx)
	resultOfQuery := adapterContextForExecQuery.Save(value)

	if resultOfQuery.Error != nil {
		return resultOfQuery.Error
	}

	return nil
}

func (adapter *DbAdapter) Destroy(ctx context.Context, options interfaces.DestroyOptions) error {
	adapterContextForExecQuery := adapter.instance(ctx)
	resultOfQuery := adapterContextForExecQuery.Table(options.Table).Delete(options.ID)

	if resultOfQuery.Error != nil {
		return resultOfQuery.Error
	}

	return nil
}
