package account

type Model struct {
	value   AccountValue
	acctype AccountType
}

func (m Model) Type() string {
	return m.acctype.String()
}

func newModel(value AccountValue, acctype AccountType) Model {
	return Model{
		value:   value,
		acctype: acctype,
	}
}
