package user

import "events-system/pkg/utils"

type UsernameVO struct {
	value string
}

func (u UsernameVO) Set(val string) error {
	if len(val) == 0 {
		return utils.GenerateError("Username.Set", "invalid value")
	}

	u.value = val

	return nil
}

func (u UsernameVO) Val() string {
	return u.value
}

func NewUsername(val string) (*UsernameVO, error) {
	u := &UsernameVO{}

	err := u.Set(val)

	if err != nil {
		return nil, err
	}

	return u, nil
}
