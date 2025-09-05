package main

import (
	"events-system/infrastructure/providers/db"
	dependency_container "events-system/pkg/di"
	"events-system/pkg/repository"
)

func InitDependencies(di *dependency_container.DependencyContainer, db *db.Database) {

	// init base repository v2
	base_repository := repository.NewBaseRepository(db)

	di.Add(
		"baseRepository",
		base_repository,
	)

}
