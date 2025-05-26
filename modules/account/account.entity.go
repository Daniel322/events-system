package account_module

import (
	"time"

	"github.com/google/uuid"
)

var ACCOUNT_TYPES = [3]string{"telegram", "vk", "gmail"}

type Account struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserId    *string   `json:"user_id"`
	AccountId *string   `json:"account_id" gorm:"unique"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
