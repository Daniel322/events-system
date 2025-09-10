package cron

import (
	"events-system/infrastructure/providers/telegram"
	"events-system/interfaces"
	"events-system/pkg/utils"
	"log"
	"os"
	"strconv"
	"time"
)

const ONE_TIME_IN_24_HOURS = 86400

type CronProvider struct {
	Name    string
	TG      *telegram.TgBotProvider
	Service interfaces.InternalUsecase
}

func NewCronProvider(TG *telegram.TgBotProvider, service interfaces.InternalUsecase) *CronProvider {
	return &CronProvider{
		Name:    "CronProvider",
		TG:      TG,
		Service: service,
	}
}

func (cron *CronProvider) Bootstrap() {
	log.Println("CRON STARTED")
	duration, err := strconv.Atoi(os.Getenv("CRON_INTERVAL"))
	if err != nil {
		utils.GenerateError(cron.Name, "invalid cron interval in env, use hardcode value")
		duration = ONE_TIME_IN_24_HOURS
	}
	utils.SetInterval(cron.TaskJob, time.Duration(duration)*time.Second)
}
