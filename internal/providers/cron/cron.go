package cron

import (
	"events-system/internal/providers/telegram"
	"events-system/internal/services"
	"events-system/internal/utils"
	"log"
	"time"
)

type CronProvider struct {
	Name         string
	TG           *telegram.TgBotProvider
	tasksService *services.TaskService
}

func NewCronProvider(TG *telegram.TgBotProvider, service *services.TaskService) *CronProvider {
	return &CronProvider{
		Name:         "CronProvider",
		TG:           TG,
		tasksService: service,
	}
}

func (cron *CronProvider) Bootstrap() {
	log.Println("CRON STARTED")
	utils.SetInterval(cron.TaskJob, time.Duration(60*60*24)*time.Second)
}
