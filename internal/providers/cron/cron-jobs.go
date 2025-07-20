package cron

import (
	"events-system/internal/utils"
	"log"
)

func (cron *CronProvider) TaskJob() error {
	log.Println("JOB STARTED")
	var tasksList, err = cron.tasksService.GetListOfTodayTasks()

	if err != nil {
		return utils.GenerateError(cron.Name, err.Error())
	}

	for _, task := range *tasksList {
		msgInfo, err := cron.tasksService.ExecTaskAndGenerateNew(task.ID.String())

		if err != nil {
			return utils.GenerateError(cron.Name, err.Error())
		}

		msg := cron.TG.NewMessage(msgInfo.ChatId)
		msg.Text = msgInfo.Text

		cron.TG.Bot.Send(msg)
	}

	return nil
}
