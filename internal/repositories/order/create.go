package order

import (
	"car-rent/internal/presentations"
	"context"
	"time"
)

func (r *repo) Create(ctx context.Context, input presentations.Order) error {

	query := `INSERT INTO orders (
		car_id,
		order_date,
		pickup_date,
		dropoff_date,
		pickup_location,
		dropoff_location,
		total_payment,
		user_id,
		created_at,
		updated_at
	) VALUES (
		:car_id,
		:order_date,
		:pickup_date,
		:dropoff_date,
		:pickup_location,
		:dropoff_location,
		:total_payment,
		:user_id,
		:created_at,
		:updated_at
	)`

	args := map[string]any{
		"car_id":           input.CarID,
		"order_date":       input.OrderDate,
		"pickup_date":      input.PickupDate,
		"dropoff_date":     input.DropoffDate,
		"pickup_location":  input.PickupLocation,
		"dropoff_location": input.DropoffLocation,
		"total_payment":    input.TotalPayment,
		"user_id":          input.UserID,
		"created_at":       time.Now().Local(),
		"updated_at":       time.Now().Local(),
	}

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return r.translateError(err)
	}

	_, err = stmt.ExecContext(ctx, args)
	if err != nil {
		return r.translateError(err)
	}
	return nil
}
