package cron

import "events-system/internal/utils"

func (cron *CronProvider) TaskJob() error {
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

		cron.TG.Bot.Send(msg)
	}

	return nil
}
