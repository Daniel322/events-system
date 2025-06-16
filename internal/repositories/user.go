package repositories

import (
	"events-system/internal/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db      *gorm.DB
	factory domain.IUserFactory
}

func NewUserRepository(db *gorm.DB, factory domain.IUserFactory) *UserRepository {
	return &UserRepository{
		db:      db,
		factory: factory,
	}
}
