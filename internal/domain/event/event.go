package event

import (
	"events-system/pkg/vo"
	"time"
)

type Model struct {
	info         vo.NonEmptyString
	date         time.Time
	notifyLevels vo.JsonField
	providers    vo.JsonField
}
