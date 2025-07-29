package presentations

import (
	"car-rent/internal/common"
	"time"
)

const (
	ErrCarsNotExist          = common.Error("err cars not exist")
	ErrCarsAlreadyExist      = common.Error("err cars already exist")
	ErrCarsAlreadyActivate   = common.Error("err cars already activate")
	ErrCarsAlreadyDeactivate = common.Error("err cars already deactivate")
)

type Cars struct {
	CarID     int       `json:"car_id" db:"car_id"`
	CarName   string    `json:"car_name" db:"car_name"`
	DayRate   float64   `json:"day_rate" db:"day_rate"`
	MonthRate float64   `json:"month_rate" db:"month_rate"`
	Image     string    `json:"image" db:"image"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAT time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
