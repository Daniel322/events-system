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
	CreateUserData
}

// func GetUsers() {
// 	var rows, err = db.Connection.Query(context.Background(), "select * from users")
// 	if (err != nil) {
// 		log.Fatal(err);
// 	}
// 	return rows.Values();
// }

func UpdateUser(data UpdateUserData) (*User, error) {
	var result User
	rowResult, err := db.Connection.Query(context.Background(), "UPDATE users SET username=$1, updated_at=NOW() WHERE id = $2", data.Username, data.Id)
	if err != nil {
		log.Fatal(err)
	}
	defer rowResult.Close()
}

func CreateUser(data CreateUserData) (*User, error) {
	var result User
	rowResult, err := db.Connection.Query(context.Background(), "INSERT INTO users (username) VALUES ($1) RETURNING *", data.Username)
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
