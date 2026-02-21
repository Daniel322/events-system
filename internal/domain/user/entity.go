package user

import (
	"encoding/json"
	"events-system/interfaces"
	"events-system/internal/components/vo"
	"events-system/internal/domain/account"
	"fmt"
	"time"
)

type Entity struct {
	interfaces.Entity
	*Model
	accounts *[]account.Entity
}

type Plain struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"username"`
	UpdatedAt time.Time `json:"created_at"`
	Username  string    `json:"updated_at"`
}

type Output struct {
	Plain
	Accounts []account.Plain `json:"accounts"`
}

func New(username vo.NonEmptyString) Entity {
	return Entity{
		interfaces.NewEntity(),
		newModel(username),
		&[]account.Entity{},
	}
}

func (e Entity) AddAccount(account account.Entity) {
	*e.accounts = append(*e.accounts, account)
}

func (e Entity) ToPlain() Plain {
	return Plain{
		ID:        e.ID.String(),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		Username:  e.Username(),
	}
}

func (e Entity) ToOutput() Output {
	plainAccs := make([]account.Plain, 0)

	if e.accounts != nil {
		for _, v := range *e.accounts {
			plainAccs = append(plainAccs, v.ToPlain())
		}
	}

	plain := Plain{
		ID:        e.ID.String(),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		Username:  e.Username(),
	}

	return Output{
		plain,
		plainAccs,
	}
}

func (e Entity) ToJSON() []byte {

	result, err := json.Marshal(e.ToOutput())

	if err != nil {
		fmt.Println(err)
	}

	return result
}
