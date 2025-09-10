package cron

import (
	"events-system/pkg/utils"
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

		// need to parse acc id to int64 for tg
		// chatId, err := strconv.ParseInt(currentAcc.AccountId, 10, 64)

		if err != nil {
			return utils.GenerateError(cron.Name, err.Error())
		}

		msg := cron.TG.NewMessage(msgInfo.ChatId)
		msg.Text = msgInfo.Text

		cron.TG.Bot.Send(msg)
	}

	return nil
}
