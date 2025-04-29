package user_module

import (
	"context"
	"errors"
	"events-system/modules/db"
	"fmt"
	"log"
	"reflect"
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

func baseQuery[T any](context context.Context, query string, args ...any) (*T, error) {
	var result T
	v := reflect.ValueOf(&result)
	fmt.Println(v.Elem().Kind())
	if v.Kind() != reflect.Ptr || (v.Elem().Kind() != reflect.Struct && v.Elem().Kind() != reflect.Slice) {
		fmt.Println("dest must be a pointer to struct")
		return nil, errors.New("dest must be a pointer to struct")
	}

	v = v.Elem()
	fields := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		fields[i] = v.Field(i).Addr().Interface()
	}

	rowResult, err := db.Connection.Query(context, query, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rowResult.Close()
	if rowResult.Next() {
		err = rowResult.Scan(
			fields...,
		)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		return nil, rowResult.Err()
	}

	return &result, nil
}

func GetUsers(options GetUserOptions) (*[]User, error) {
	var query = "SELECT * FROM users LIMIT = $1 OFFSET = $2"
	result, err := baseQuery[[]User](context.Background(), query, options.Limit, options.Skip)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

func UpdateUser(data UpdateUserData) (*User, error) {
	var query = "UPDATE users SET username=$1, updated_at=NOW() WHERE id = $2 RETURNING *"
	result, err := baseQuery[User](context.Background(), query, data.Username, data.Id)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

func CreateUser(data CreateUserData) (*User, error) {
	var query = "INSERT INTO users (username) VALUES ($1) RETURNING *"
	result, err := baseQuery[User](context.Background(), query, data.Username)
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
