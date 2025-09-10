package cron

import (
	"events-system/pkg/utils"
	"log"
	"strconv"
)

func (cron *CronProvider) TaskJob() error {
	log.Println("JOB STARTED")
	var tasksList, err = cron.Service.GetListOfTodayTasks()

	if err != nil {
		return utils.GenerateError(cron.Name, err.Error())
	}

	for _, task := range *tasksList {
		msgInfo, err := cron.Service.ExecTask(task.ID.String())

		if err != nil {
			return utils.GenerateError(cron.Name, err.Error())
		}

		// need to parse acc id to int64 for tg
		chatId, err := strconv.ParseInt(msgInfo.ChatId, 10, 64)

		if err != nil {
			return utils.GenerateError(cron.Name, err.Error())
		}

		msg := cron.TG.NewMessage(chatId)
		msg.Text = msgInfo.Text

		cron.TG.Bot.Send(msg)
	}

	return nil
}
