package account

type AccountType int

const (
	Telegram AccountType = iota
	VK
	Gmail
)

var accountTypes = map[AccountType]string{
	Telegram: "telegram",
	VK:       "vk",
	Gmail:    "gmail",
}

func (val AccountType) String() string {
	return accountTypes[val]
}

type Account struct {
	UserId    string      `json:"user_id"`
	AccountId string      `json:"account_id"`
	Type      AccountType `json:"type"`
}
