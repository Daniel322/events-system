package task

import "events-system/pkg/utils"

type TaskProvider int

const (
	Telegram TaskProvider = iota
	Mail
)

var TASK_PROVIDERS = map[TaskProvider]string{
	Telegram: "telegram",
	Mail:     "mail",
}

func (task TaskProvider) String() string {
	return TASK_PROVIDERS[task]
}

func NewTaskProvider(s string) (TaskProvider, error) {
	switch s {
	case "telegram":
		return TaskProvider(0), nil
	case "mail":
		return TaskProvider(1), nil
	default:
		return TaskProvider(-1), utils.GenerateError("TaskProvider", "invalid task provider")
	}
}
