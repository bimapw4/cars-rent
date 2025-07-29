package entity

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Order struct {
	CarID           int       `json:"car_id"`
	OrderDate       time.Time `json:"order_date"`
	PickupDate      time.Time `json:"pickup_date"`
	DropoffDate     time.Time `json:"dropoff_date"`
	PickupLocation  string    `json:"pickup_location"`
	DropoffLocation string    `json:"dropoff_location"`
}

func (o Order) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.CarID, validation.Required, validation.Min(1)),
		validation.Field(&o.OrderDate, validation.Required),
		validation.Field(&o.PickupDate, validation.Required),
		validation.Field(&o.DropoffDate, validation.Required, validation.By(func(value interface{}) error {
			if o.PickupDate.After(o.DropoffDate) {
				return validation.NewError("validation_dropoff", "dropoff_date must be after pickup_date")
			}
			return nil
		})),
		validation.Field(&o.PickupLocation, validation.Required, validation.Length(2, 100)),
		validation.Field(&o.DropoffLocation, validation.Required, validation.Length(2, 100)),
	)
}
