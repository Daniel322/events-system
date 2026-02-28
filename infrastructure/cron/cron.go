package cron

import (
	"events-system/infrastructure/config"
	"events-system/infrastructure/cron/jobs"
	"events-system/interfaces"
	"events-system/pkg/utils"
	"log"
	"os"
	"strconv"
	"time"
)

const ONE_TIME_IN_24_HOURS = 86400

type CronProvider struct {
	Name   string
	logger *log.Logger
	stops  *[]chan bool
	sender interfaces.Sender
}

func NewCronProvider(sender interfaces.Sender) *CronProvider {
	var logger = log.New(os.Stdout, "CronProvider"+" ", log.LstdFlags)
	return &CronProvider{
		Name:   "CronProvider",
		logger: logger,
		stops:  new([]chan bool),
		sender: sender,
	}
}

func (cron *CronProvider) Bootstrap() {
	cron.logger.Println("BOOTSTRAP STARTED")
	strDuration, err := config.Config.CRON_INTERVAL()

	if err != nil {
		utils.GenerateError(cron.Name, "error on get cron interval from config")
	}

	duration, err := strconv.Atoi(strDuration)
	if err != nil {
		utils.GenerateError(cron.Name, "invalid cron interval in env, use hardcode value")
		duration = ONE_TIME_IN_24_HOURS
	}

	taskJob := jobs.NewTaskJob(cron.sender)

	stop := utils.SetInterval(taskJob.Run, time.Duration(duration)*time.Second)
	*cron.stops = append(*cron.stops, stop)
}

func (cron *CronProvider) Stop() {
	for _, stop := range *cron.stops {
		cron.logger.Println("JOB STOP BY STOP CHANNEL")
		stop <- true
	}
	cron.logger.Println("CRON PROVIDER STOPPED")
}
