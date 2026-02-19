package task

import "events-system/pkg/utils"

type TaskType int

var SUPPORTED_TYPES = []string{"month", "week", "tomorrow", "today"}

const (
	Today TaskType = iota
	Tomorrow
	Week
	Month
)

var TASK_TYPES = map[TaskType]string{
	Today:    "today",
	Tomorrow: "tomorrow",
	Week:     "week",
	Month:    "month",
}

func (task TaskType) String() string {
	return TASK_TYPES[task]
}

func NewTaskType(s string) (TaskType, error) {
	switch s {
	case "today":
		return TaskType(0), nil
	case "tomorrow":
		return TaskType(1), nil
	case "week":
		return TaskType(2), nil
	case "month":
		return TaskType(3), nil
	default:
		return TaskType(-1), utils.GenerateError("TaskType", "invalid task type")
	}
}

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
