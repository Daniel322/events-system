package event

import (
	"events-system/pkg/vo"
	"time"
)

type Model struct {
	info         vo.NonEmptyString
	date         time.Time
	Type         EventType
	notifyLevels NotifyLevels
	providers    vo.JsonField // mail, telegram
}
