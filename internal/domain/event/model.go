package event

import (
	"events-system/pkg/vo"
	"time"
)

type Model struct {
	info         vo.NonEmptyString
	date         time.Time
	Type         string
	notifyLevels vo.JsonField
	providers    vo.JsonField
}
