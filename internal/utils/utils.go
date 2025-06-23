package utils

import (
	"fmt"

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

func ParseId[T ParsedTypes](id T, flow ParseFlow) (uuid.UUID, string, error) {
	var (
		uuidVersion   uuid.UUID
		stringVersion string
		err           error
	)

	switch flow {
	case String:
		if str, ok := any(id).(string); ok {
			stringVersion = str
			uuidVersion, err = uuid.Parse(str)
		} else {
			err = fmt.Errorf("expected string, got %T", id)
		}
	case UUID:
		if u, ok := any(id).(uuid.UUID); ok {
			uuidVersion = u
			stringVersion = u.String()
		} else {
			err = fmt.Errorf("expected uuid.UUID, got %T", id)
		}
	}

	return uuidVersion, stringVersion, err
}
