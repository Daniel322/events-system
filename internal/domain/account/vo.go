package account

import (
	"events-system/pkg/utils"
	"events-system/pkg/vo"
	"net/mail"
)

type AccountValue struct {
	vo.NonEmptyString
}

type AccountType int

const (
	Telegram AccountType = iota
	Mail
)

var SUPPORTED_ACCOUNT_TYPES = map[AccountType]string{
	Telegram: "telegram",
	Mail:     "mail",
}

func (acc AccountType) String() string {
	return SUPPORTED_ACCOUNT_TYPES[acc]
}

func NewAccountType(s string) (AccountType, error) {
	switch s {
	case "telegram":
		return AccountType(0), nil
	case "mail":
		return AccountType(1), nil
	default:
		return AccountType(-1), utils.GenerateError("AccountType", "invalid acc type")
	}
}

func NewAccountValue(s string, t AccountType) (AccountValue, error) {
	if t == 0 {
		if ok := IsEmail(s); !ok {
			return AccountValue{}, utils.GenerateError("AccountValue", "invalid value for type mail")
		}
	}

	res, err := vo.NewNonEmptyString(s)
	return AccountValue{res}, err
}

func IsEmail(email string) bool {
	addr, err := mail.ParseAddress(email)
	return err == nil && addr.Address == email
}
