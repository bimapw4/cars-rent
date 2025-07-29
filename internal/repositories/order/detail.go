package order

import (
	"car-rent/internal/presentations"
	"context"
)

func (r *repo) Detail(ctx context.Context, orderID int) (*presentations.Order, error) {

	var (
		result presentations.Order
	)
	query := `select * from orders where order_id=:order_id and is_active=true`

	args := map[string]any{
		"order_id": orderID,
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

func (r *repo) DetailWithoutIsActive(ctx context.Context, orderID int) (*presentations.Order, error) {

	var (
		result presentations.Order
	)
	query := `select * from orders where order_id=:order_id`

	args := map[string]any{
		"order_id": orderID,
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

func (r *repo) Latest(ctx context.Context) (*presentations.Order, error) {

	var (
		result presentations.Order
	)
	query := `select * from orders order by order_id desc limit 1`

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
