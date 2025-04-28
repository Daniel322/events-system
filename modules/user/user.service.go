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

func baseQuery(context context.Context, query string, args ...any) (*User, error) {
	var result User
	rowResult, err := db.Connection.Query(context, query, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rowResult.Close()
	if rowResult.Next() {
		err = rowResult.Scan(
			&result.Id,
			&result.Username,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		return nil, rowResult.Err()
	}

	return &result, nil
}

func UpdateUser(data UpdateUserData) (*User, error) {
	var query = "UPDATE users SET username=$1, updated_at=NOW() WHERE id = $2 RETURNING *"
	result, err := baseQuery(context.Background(), query, data.Username, data.Id)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

func CreateUser(data CreateUserData) (*User, error) {
	var query = "INSERT INTO users (username) VALUES ($1) RETURNING *"
	result, err := baseQuery(context.Background(), query, data.Username)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}
