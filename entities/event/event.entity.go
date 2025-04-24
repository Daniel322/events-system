package event

type Event struct {
	Id           string   `json:"id"`
	UserId       string   `json:"user_id"`
	Info         string   `json:"info"`
	NotifyLevels []string `json:"notify_levels"`
	Providers    []string `json:"providers"`
}
