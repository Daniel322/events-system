package account

import "events-system/pkg/vo"

type Model struct {
	value   vo.NonEmptyString
	acctype vo.AccountType
}

func (m Model) Type() string {
	return m.acctype.String()
}

func newModel(value vo.NonEmptyString, acctype vo.AccountType) Model {
	return Model{
		value:   value,
		acctype: acctype,
	}
}
