package entities

import (
	"database/sql/driver"
	"encoding/json"
	"events-system/pkg/utils"
)

type JsonField []string

func (value *JsonField) Scan(src interface{}) error {
	if src == nil {
		*value = nil
		return nil
	}

	var bytes []byte
	switch v := src.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		// Если это не []byte и не string, попробуем замаршалить
		var err error
		bytes, err = json.Marshal(src)
		if err != nil {
			return utils.GenerateError("json field scan", "src value cannot be converted to []byte")
		}
	}

	err := json.Unmarshal(bytes, value)
	if err != nil {
		return utils.GenerateError("json field scan", "failed to unmarshal JSON: "+err.Error())
	}

	return nil
}

func (value JsonField) Value() (driver.Value, error) {
	if len(value) == 0 {
		return nil, nil
	}
	return json.Marshal(value)
}
