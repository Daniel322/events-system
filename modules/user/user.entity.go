package user_module

import "database/sql"

type User struct {
	Id        string       `json:"id"`
	Username  string       `json:"username"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
