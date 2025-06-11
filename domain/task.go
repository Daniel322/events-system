package domain

type Task struct {
	Id        string `json:"id"`
	EventId   string `json:"event_id"`
	AccountId string `json:"account_id"`
	Date      string `json:"date"`
}
