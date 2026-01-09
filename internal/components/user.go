package components

import (
	"events-system/interfaces"
	entities "events-system/internal/entity"
)

type UserComponent struct {
	interfaces.RepositoryV2
	entities.User
}
