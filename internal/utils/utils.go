package utils

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type ParseFlow string

type ParsedTypes interface {
	string | uuid.UUID
}

const (
	String ParseFlow = "string"
	UUID   ParseFlow = "uuid"
)

func ParseId[T ParsedTypes](id T) (uuid.UUID, string, error) {
	var (
		uuidVersion   uuid.UUID
		stringVersion string
		err           error
	)

	if str, ok := any(id).(string); ok {
		stringVersion = str
		uuidVersion, err = uuid.Parse(str)
	} else {
		if u, ok := any(id).(uuid.UUID); ok {
			uuidVersion = u
			stringVersion = u.String()
		} else {
			err = fmt.Errorf("invalid type, got %T", id)
		}
	}

	return uuidVersion, stringVersion, err
}

func GenerateError(Name string, Message string) error {
	log.SetPrefix("ERROR " + Name + " ")
	log.Println(" " + Message)
	return errors.New("Error in " + Name + ": " + Message)
}

func SetInterval(callback func(), interval time.Duration) chan bool {
	ticker := time.NewTicker(interval)
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				callback()
			case <-stop:
				ticker.Stop()
				return
			}
		}
	}()

	return stop
}
