package cars

import (
	"car-rent/internal/presentations"
	"context"
	"time"
)

func (r *repo) Update(ctx context.Context, payload *presentations.Cars) error {

	query := `
		update 
			cars 
		set 
			car_name=:car_name,
			day_rate=:day_rate,
			month_rate=:month_rate, 
			image=:image, 
			updated_at=:updated_at
		where 
			car_id=:car_id`

	args := map[string]any{
		"car_id":     payload.CarID,
		"car_name":   payload.CarName,
		"day_rate":   payload.DayRate,
		"month_rate": payload.MonthRate,
		"image":      payload.Image,
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

func (r *repo) UpdateIsActive(ctx context.Context, carsID int, isActive bool) error {

	query := `
		update 
			cars 
		set 
			is_active=:is_active, 
			updated_at=:updated_at
		where 
			car_id=:car_id`

	args := map[string]any{
		"car_id":     carsID,
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
