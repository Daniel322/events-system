package cron

import (
	"events-system/infrastructure/config"
	"events-system/infrastructure/cron/jobs"
	"events-system/internal/application/commands"
	"events-system/internal/application/queries"
	"events-system/pkg/utils"
	"log"
	"os"
	"strconv"
	"time"
)

const ONE_TIME_IN_24_HOURS = 86400

type CronProvider struct {
	Name            string
	Logger          *log.Logger
	Stops           *[]chan bool
	TasksListAction *queries.TasksList
	ExecTaskCmd     *commands.Exectask
}

func NewCronProvider(action *queries.TasksList, cmd *commands.Exectask) *CronProvider {
	var logger = log.New(os.Stdout, "CronProvider"+" ", log.LstdFlags)
	return &CronProvider{
		Name:            "CronProvider",
		Logger:          logger,
		Stops:           new([]chan bool),
		TasksListAction: action,
		ExecTaskCmd:     cmd,
	}
}

func (cron *CronProvider) Bootstrap() {
	cron.Logger.Println("BOOTSTRAP STARTED")
	strDuration, err := config.Config.CRON_INTERVAL()

	if err != nil {
		utils.GenerateError(cron.Name, "error on get cron interval from config")
	}

	duration, err := strconv.Atoi(strDuration)
	if err != nil {
		utils.GenerateError(cron.Name, "invalid cron interval in env, use hardcode value")
		duration = ONE_TIME_IN_24_HOURS
	}

	taskJob := jobs.NewTaskJob(cron.TasksListAction, cron.ExecTaskCmd)

	stop := utils.SetInterval(taskJob.Run, time.Duration(duration)*time.Second)
	*cron.Stops = append(*cron.Stops, stop)
}

func (cron *CronProvider) Stop() {
	for _, stop := range *cron.Stops {
		cron.Logger.Println("JOB STOP BY STOP CHANNEL")
		stop <- true
	}
	cron.Logger.Println("CRON PROVIDER STOPPED")
}
