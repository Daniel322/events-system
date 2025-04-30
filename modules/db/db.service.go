package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/jackc/pgx/v5"
)

var err error
var Connection *pgx.Conn

func ConnectDatabase(context context.Context) {
	Connection, err = pgx.Connect(context, os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		log.Fatal(err)
	}
}

func CheckConnection() bool {
	return !Connection.IsClosed()
}

func Close(context context.Context) {
	Connection.Close(context)
}

func BaseQuery[T any](context context.Context, query string, args ...any) (*T, error) {
	var result T
	v := reflect.ValueOf(&result)
	if v.Kind() != reflect.Ptr || (v.Elem().Kind() != reflect.Struct) {
		fmt.Println("dest must be a pointer to struct")
		return nil, errors.New("dest must be a pointer to struct")
	}

	v = v.Elem()
	fields := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		fields[i] = v.Field(i).Addr().Interface()
	}

	rowResult, err := Connection.Query(context, query, args...)
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
