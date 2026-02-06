package vo

import "events-system/pkg/utils"

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
