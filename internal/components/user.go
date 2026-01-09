package components

import (
	"events-system/interfaces"
	"time"

	"github.com/google/uuid"
)

type UserComponent struct {
	interfaces.RepositoryV2
	ID        uuid.UUID
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
