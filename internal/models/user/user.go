package user

import (
	"events-system/internal/models/account"
	"events-system/internal/models/event"
)

type Model struct {
	username string
	accounts []account.Model
	events   []event.Model
}
