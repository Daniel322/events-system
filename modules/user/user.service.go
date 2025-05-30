package user_module

import (
	"context"
	"events-system/modules/db"
	"log"
)

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

type UserService interface {
	GetUsers(options GetUserOptions) (*[]User, error)
	GetUserById(id string) (*User, error)
	UpdateUser(data UpdateUserData, operationContext context.Context) (*User, error)
	CreateUser(data CreateUserData, operationContext context.Context) (*User, error)
	DeleteUser(id string, operationContext context.Context) (bool, error)
}

func GetUsers(options GetUserOptions) (*[]User, error) {
	var users []User

	result := db.Connection.Limit(options.Limit).Offset(options.Skip).Table("users").Find(&users)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return &users, nil
}

func GetUserById(id string) (*User, error) {
	var user User

	result := db.Connection.Where("id = ?", id).First(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func UpdateUser(data UpdateUserData, operationContext context.Context) (*User, error) {
	var user User
	result := db.Connection.Model(&user).Where("id = ?", data.Id).Update("username", data.Username)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func CreateUser(data CreateUserData, operationContext context.Context) (*User, error) {
	user := User{Username: data.Username}

	result := db.Connection.Create(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func DeleteUser(id string, operationContext context.Context) (bool, error) {
	result := db.Connection.Table("users").Delete(&User{}, id)

	if result.Error != nil {
		log.Fatal(result.Error)
		return false, result.Error
	}

	return true, nil
}
