package presentations

import (
	"car-rent/internal/common"
	"time"
)

const (
	ErrUserNotExist     = common.Error("err users not exist")
	ErrUserAlreadyExist = common.Error("err users already exist")
)

type Users struct {
	UserID    int       `json:"user_id" db:"user_id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password,omitempty" db:"password"`
	IsAdmin   bool      `json:"is_admin" db:"is_admin"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
