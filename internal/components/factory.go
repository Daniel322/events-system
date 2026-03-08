package components

import (
	"events-system/interfaces"
	"log"
	"os"
)

type Factory struct {
	Entity     string
	Repository interfaces.Repository
	Logger     *log.Logger
}

func NewFactory(entity string, repo interfaces.Repository) *Factory {
	var logger = log.New(os.Stdout, entity+" factory ", log.LstdFlags)

	return &Factory{
		Logger:     logger,
		Entity:     entity,
		Repository: repo,
	}
}
