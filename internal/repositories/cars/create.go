package cars

import (
	"car-rent/internal/presentations"
	"context"
	"time"
)

func (r *repo) Create(ctx context.Context, input presentations.Cars) error {

	query := `INSERT INTO cars (
		car_name,
		day_rate,
		month_rate,
		image,
		is_active,
		created_at,
		updated_at
	) VALUES (
		:car_name,
		:day_rate,
		:month_rate,
		:image,
		:is_active,
		:created_at,
		:updated_at
	)`

	args := map[string]any{
		"car_name":   input.CarName,
		"day_rate":   input.DayRate,
		"month_rate": input.MonthRate,
		"image":      input.Image,
		"is_active":  input.IsActive,
		"created_at": time.Now().Local(),
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
