package user

import (
	"encoding/json"
	"events-system/interfaces"
	"events-system/pkg/vo"
	"fmt"
	"time"
)

type Entity struct {
	interfaces.Entity
	*Model
}

type Plain struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"username"`
	UpdatedAt time.Time `json:"createdAt"`
	Username  string    `json:"updatedAt"`
}

func New(username vo.NonEmptyString) Entity {
	return Entity{
		interfaces.NewEntity(),
		newModel(username),
	}
}

func (e Entity) ToPlain() Plain {
	return Plain{
		ID:        e.ID.String(),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		Username:  e.Username(),
	}
}

func (e Entity) ToJSON() []byte {

	result, err := json.Marshal(e.ToPlain())

	if err != nil {
		fmt.Println(err)
	}

	return result
}
