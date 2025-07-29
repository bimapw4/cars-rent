package presentations

import (
	"car-rent/internal/common"
	"time"
)

const (
	ErrOrdersNotExist          = common.Error("err orders not exist")
	ErrOrdersAlreadyExist      = common.Error("err orders already exist")
	ErrOrdersAlreadyActivate   = common.Error("err orders already activate")
	ErrOrdersAlreadyDeactivate = common.Error("err orders already deactivate")
)

type Order struct {
	OrderID         int       `json:"order_id" db:"order_id"`
	CarID           int       `json:"car_id" db:"car_id"`
	OrderDate       time.Time `json:"order_date" db:"order_date"`
	PickupDate      time.Time `json:"pickup_date" db:"pickup_date"`
	DropoffDate     time.Time `json:"dropoff_date" db:"dropoff_date"`
	PickupLocation  string    `json:"pickup_location" db:"pickup_location"`
	DropoffLocation string    `json:"dropoff_location" db:"dropoff_location"`
	TotalPayment    float64   `json:"total_payment" db:"total_payment"`
	UserID          int       `json:"user_id" db:"user_id"`
	IsActive        bool      `json:"is_active" db:"is_active"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}
