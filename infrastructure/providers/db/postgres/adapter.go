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

func (adapter *DbAdapter) Save(ctx *context.Context, value interface{}) error {
	adapter.Instance.Save(value)

	return nil
}
