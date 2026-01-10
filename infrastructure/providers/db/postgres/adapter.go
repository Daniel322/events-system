package pg_db

import (
	"context"
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

func (adapter *DbAdapter) Save(ctx context.Context, value interface{}) error {
	if ctx.Value("transaction") != nil {
		res := ctx.Value("transaction").(*gorm.DB).Save(value)

		if res.Error != nil {
			return res.Error
		}

		return nil
	}
	res := adapter.Instance.WithContext(ctx).Save(value)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
