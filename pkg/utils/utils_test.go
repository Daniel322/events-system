package utils

import (
	"testing"

	"github.com/google/uuid"
)

func TestParseId(t *testing.T) {
	testUuid := uuid.New()

	_, _, err := ParseId(testUuid)

	if err != nil {
		t.Error("ParseId should return uuid in string format, but return error")
	}

	invalidTestUuid := "invalidstring"

	_, _, err = ParseId(invalidTestUuid)

	if err == nil {
		t.Error("ParseId should return error, but error was nil")
	}
}

func TestGenerateError(t *testing.T) {
	err := GenerateError("testName", "some error")

	if err == nil {
		t.Error("should be error, but get nil")
	}
}
