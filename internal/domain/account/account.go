package account

import "events-system/pkg/vo"

type Model struct {
	value   vo.NonEmptyString
	acctype vo.NonEmptyString
}
