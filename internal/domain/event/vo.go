package event

import (
	"events-system/pkg/utils"
	"events-system/pkg/vo"
	"slices"
)

// Notify levels

type NotifyLevels vo.JsonField

var SUPPORTED_NOTIFY_LEVELS = []string{"today", "tomorrow", "month", "week"}

func NewNotifyLevels(values []string) (vo.JsonField, error) {
	result := vo.JsonField{}
	for _, v := range values {
		if ok := slices.Contains(SUPPORTED_NOTIFY_LEVELS, v); !ok {
			return vo.JsonField{}, utils.GenerateError("NotifyLevels", "invalid notify level")
		} else {
			result = append(result, v)
		}
	}

	return result, nil
}

// Providers

type Providers vo.JsonField

var SUPPORTED_PROVIDERS = []string{"mail", "telegram"}

func NewProviders(values []string) (vo.JsonField, error) {
	result := vo.JsonField{}
	for _, v := range values {
		if ok := slices.Contains(SUPPORTED_PROVIDERS, v); !ok {
			return vo.JsonField{}, utils.GenerateError("Providers", "invalid provider")
		} else {
			result = append(result, v)
		}
	}

	return result, nil
}
