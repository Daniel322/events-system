package components

import (
	"events-system/interfaces"
	"log"
	"os"
)

type Factory struct {
	Entity     string
	Repository interfaces.RepositoryV2
	Logger     *log.Logger
}

func NewFactory(entity string, repo interfaces.RepositoryV2) *Factory {
	var logger = log.New(os.Stdout, entity+" factory ", log.LstdFlags)

	return &Factory{
		Logger:     logger,
		Entity:     entity,
		Repository: repo,
	}
}
