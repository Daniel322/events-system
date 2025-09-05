package entities

import (
	"database/sql/driver"
	"encoding/json"
	"events-system/pkg/utils"
)

type NotifyLevel []string

func (value *NotifyLevel) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return utils.GenerateError("Notify level scan", "src value cannot be cast to []byte")
	}
	return json.Unmarshal(bytes, value)
}

func (value NotifyLevel) Value() (driver.Value, error) {
	if len(value) == 0 {
		return nil, nil
	}
	return json.Marshal(value)
}
