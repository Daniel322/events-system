package entities

import (
	"database/sql/driver"
	"encoding/json"
	"events-system/pkg/utils"
)

type JsonField []string

func (value *JsonField) Scan(src interface{}) error {
	bytes, ok := json.Marshal(src)
	if ok != nil {
		return utils.GenerateError("json field scan", "src value cannot be cast to []byte")
	}
	return json.Unmarshal(bytes, value)
}

func (value JsonField) Value() (driver.Value, error) {
	if len(value) == 0 {
		return nil, nil
	}
	return json.Marshal(value)
}
