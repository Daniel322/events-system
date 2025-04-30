package user_module

import (
	"context"
	"events-system/modules/db"
	"log"
)

type UserService struct {
}

type CreateUserData struct {
	Username string `json:"username"`
}

type UpdateUserData struct {
	Id string `json:"id"`
	// TODO: find how can inherit upper struct
	Username string `json:"username"`
	// CreateUserData
}

type GetUserOptions struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}

func GetUsers(options GetUserOptions) (*[]User, error) {
	var query = "SELECT * FROM users LIMIT $1 OFFSET $2"
	rows, err := db.Connection.Query(context.Background(), query, options.Limit, options.Skip)
	if err != nil {
		log.Fatal(err)
	}

	var result []User

	for rows.Next() {
		var iterationScanValue User
		err = rows.Scan(
			&iterationScanValue.Id,
			&iterationScanValue.Username,
			&iterationScanValue.CreatedAt,
			&iterationScanValue.UpdatedAt,
		)
		if err != nil {
			log.Fatal(err)
		} else {
			result = append(
				result,
				User{
					Id:        string(iterationScanValue.Id),
					Username:  iterationScanValue.Username,
					CreatedAt: iterationScanValue.CreatedAt,
					UpdatedAt: iterationScanValue.UpdatedAt,
				},
			)
		}
	}

	return &result, err
}

func GetUserById(id string) (*User, error) {
	var query = "SELECT * FROM users WHERE id = $1"
	result, err := db.BaseQuery[User](context.Background(), query, id)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

func UpdateUser(data UpdateUserData) (*User, error) {
	var query = "UPDATE users SET username=$1, updated_at=NOW() WHERE id = $2 RETURNING *"
	result, err := db.BaseQuery[User](context.Background(), query, data.Username, data.Id)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

func CreateUser(data CreateUserData) (*User, error) {
	var query = "INSERT INTO users (username) VALUES ($1) RETURNING *"
	result, err := db.BaseQuery[User](context.Background(), query, data.Username)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

func DeleteUser(id string) (bool, error) {
	query := "DELETE FROM users WHERE id = $1"
	_, err := db.Connection.Exec(context.Background(), query, id)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, err
}
