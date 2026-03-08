package task

import (
	"time"
)

type Model struct {
	date     time.Time
	taskType TaskType
	provider TaskProvider
}

func newModel(
	date time.Time,
	taskType TaskType,
	provider TaskProvider,
) Model {
	return Model{
		taskType: taskType,
		date:     date,
		provider: provider,
	}
}
