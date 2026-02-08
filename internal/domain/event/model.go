package event

import (
	"events-system/pkg/vo"
	"time"
)

type Model struct {
	info         vo.NonEmptyString
	date         time.Time
	Type         string       // hb, reminder?
	notifyLevels vo.JsonField // month, week, tomorrow, today
	providers    vo.JsonField // mail, telegram
}
