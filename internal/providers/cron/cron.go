package cron

import (
	"events-system/internal/providers/telegram"
	"events-system/internal/services"
	"events-system/internal/utils"
	"log"
	"os"
	"strconv"
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
	duration, err := strconv.Atoi(os.Getenv("CRON_INTERVAL"))
	if err != nil {
		utils.GenerateError(cron.Name, "invalid cron interval in env, use hardcode value")
		duration = 86400
	}
	utils.SetInterval(cron.TaskJob, time.Duration(duration)*time.Second)
}
