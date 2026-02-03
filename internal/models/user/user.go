package user

import (
	"events-system/internal/models/account"
	"events-system/internal/models/event"
	"events-system/pkg/vo"
)

type Model struct {
	username vo.NonEmptyString
	accounts []account.Model
	events   []event.Model
}
