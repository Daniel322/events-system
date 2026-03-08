package pg_db

import (
	"context"
	"events-system/interfaces"
	"events-system/pkg/utils"
	"log"
	"os"

	"gorm.io/gorm"
)

type DbAdapter struct {
	Instance *gorm.DB
	Logger   *log.Logger
}

const NAME = "DB Adapter"

var Adapter *DbAdapter

func InitAdapter(instance *gorm.DB) {
	var logger = log.New(os.Stdout, NAME+" ", log.LstdFlags)

	Adapter = &DbAdapter{
		Instance: instance,
		Logger:   logger,
	}
}

func (adapter *DbAdapter) CreateTransaction(ctx context.Context) context.Context {
	if ctx.Value("transaction") == nil {
		tx := adapter.Instance.Begin()
		ctx = context.WithValue(ctx, "transaction", tx)
	}

	return ctx
}

func (adapter *DbAdapter) Commit(ctx context.Context) {
	if ctx.Value("transaction") != nil {
		ctx.Value("transaction").(*gorm.DB).Commit()
	}
}

func (adapter *DbAdapter) Rollback(ctx context.Context) {
	if ctx.Value("transaction") != nil {
		ctx.Value("transaction").(*gorm.DB).Rollback()
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
	resultOfQuery := adapterContextForExecQuery.Table(ctx.Value("tableName").(string)).Save(value)

	if resultOfQuery.Error != nil {
		return utils.GenerateError(NAME, resultOfQuery.Error.Error())
	}

	return nil
}

func (adapter *DbAdapter) Destroy(ctx context.Context, options interfaces.DestroyOptions) error {
	adapterContextForExecQuery := adapter.instance(ctx)
	resultOfQuery := adapterContextForExecQuery.Table(options.Table).Where("id = ?", options.ID).Delete(options.ID)

	if resultOfQuery.Error != nil {
		return utils.GenerateError(NAME, resultOfQuery.Error.Error())
	}

	return nil
}

func (adapter *DbAdapter) Find(
	ctx context.Context,
	options map[string]interface{},
) error {
	ptr := ctx.Value("ptr")

	result := adapter.instance(ctx).Table(ctx.Value("tableName").(string)).Find(ptr, options)

	if result.Error != nil {
		return utils.GenerateError(NAME, result.Error.Error())
	}

	return nil
}
