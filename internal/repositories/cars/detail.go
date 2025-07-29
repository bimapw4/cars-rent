package cars

import (
	"car-rent/internal/presentations"
	"context"
)

func (r *repo) Detail(ctx context.Context, carsID int) (*presentations.Cars, error) {

	var (
		result presentations.Cars
	)
	query := `select * from cars where car_id=:car_id and is_active=true`

	args := map[string]any{
		"car_id": carsID,
	}

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, r.translateError(err)
	}

	err = stmt.GetContext(ctx, &result, args)
	if err != nil {
		return nil, r.translateError(err)
	}
	return &result, nil
}

func (r *repo) Latest(ctx context.Context) (*presentations.Cars, error) {

	var (
		result presentations.Cars
	)
	query := `select * from cars order by car_id desc limit 1`

	args := map[string]any{}

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, r.translateError(err)
	}

	err = stmt.GetContext(ctx, &result, args)
	if err != nil {
		return nil, r.translateError(err)
	}
	return &result, nil
}
