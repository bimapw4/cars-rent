package order

import (
	"car-rent/internal/presentations"
	"context"
	"time"
)

func (r *repo) Update(ctx context.Context, input *presentations.Order) error {

	query := `
		UPDATE orders
		SET
			car_id = :car_id,
			order_date = :order_date,
			pickup_date = :pickup_date,
			dropoff_date = :dropoff_date,
			pickup_location = :pickup_location,
			total_payment = :total_payment,
			dropoff_location = :dropoff_location,
			updated_at = :updated_at
		WHERE
			order_id = :order_id`

	args := map[string]any{
		"order_id":         input.OrderID,
		"car_id":           input.CarID,
		"order_date":       input.OrderDate,
		"pickup_date":      input.PickupDate,
		"dropoff_date":     input.DropoffDate,
		"pickup_location":  input.PickupLocation,
		"dropoff_location": input.DropoffLocation,
		"total_payment":    input.TotalPayment,
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

func (r *repo) UpdateIsActive(ctx context.Context, orderID int, isActive bool) error {

	query := `
		update 
			orders 
		set 
			is_active=:is_active, 
			updated_at=:updated_at
		where 
			car_id=:car_id`

	args := map[string]any{
		"car_id":     orderID,
		"is_active":  isActive,
		"updated_at": time.Now().Local(),
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
