package vo

import (
	"events-system/pkg/utils"
	"strings"
)

type NonEmptyString struct {
	value string
}

func NewNonEmptyString(s string) (NonEmptyString, error) {
	// Optional: Trim leading/trailing whitespace before validation if desired
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return NonEmptyString{}, utils.GenerateError("NonEmptyString", "string cannot be empty or contain only whitespace")
	}
	return NonEmptyString{value: trimmed}, nil
}

// Get returns the underlying string value.
func (n NonEmptyString) Get() string {
	return n.value
}
