package event_module

import (
	"database/sql"
	"time"
)

type Event struct {
	Id           string       `json:"id"`
	UserId       string       `json:"user_id"`
	Info         string       `json:"info"`
	Date         time.Time    `json:"date"`
	NotifyLevels []string     `json:"notify_levels"`
	Providers    []byte       `json:"providers"`
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
}
