package main

import (
	"events-system/infrastructure/providers/db"
	"events-system/internal/services"
	dependency_container "events-system/pkg/di"
	"events-system/pkg/repository"
)

func InitDependencies(di *dependency_container.DependencyContainer, db *db.Database) {

	// init base repository v2
	base_repository := repository.NewBaseRepository(db)

	user_service := services.NewUserService(base_repository)

	di.Add(
		"baseRepository",
		base_repository,
	)

	di.Add(
		"userService",
		user_service,
	)

}
