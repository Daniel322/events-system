package repository

type ModelName int

const (
	Users ModelName = iota
	Accounts
	Events
	Tasks
)

var modelNames = map[ModelName]string{
	Users:    "users",
	Accounts: "accounts",
	Events:   "events",
	Tasks:    "tasks",
}

func (mod ModelName) String() string {
	return modelNames[mod]
}
