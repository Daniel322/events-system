package task

import (
	"events-system/pkg/vo"
	"time"
)

type Model struct {
	date     time.Time
	taskType vo.EventType
	provider TaskProvider
}

func newModel(
	date time.Time,
	taskType vo.EventType,
	provider TaskProvider,
) Model {
	return Model{
		taskType: taskType,
		date:     date,
		provider: provider,
	}
}
