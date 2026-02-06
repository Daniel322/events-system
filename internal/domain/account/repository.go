package account

import (
	"events-system/interfaces"
	"events-system/internal/components"
)

type AccRepo struct {
	components.Factory
}

func NewAccRepo(repo interfaces.RepositoryV2) *AccRepo {
	return &AccRepo{
		Factory: *components.NewFactory("Account", repo),
	}
}
