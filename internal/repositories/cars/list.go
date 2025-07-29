package cars

import (
	"car-rent/internal/presentations"
	"car-rent/pkg/meta"
	"context"
	"fmt"
	"strings"
	"time"
)

func (r *repo) List(ctx context.Context, m *meta.Params) ([]presentations.Cars, error) {
	var (
		result = []presentations.Cars{}
	)
	q, err := meta.Parse(m)
	if err != nil {
		return nil, err
	}
	query := `SELECT* FROM cars where 1=1 and is_active=true ORDER BY created_at DESC OFFSET :offset LIMIT :limit`

	query = strings.Replace(
		query,
		" ORDER BY created_at DESC ",
		fmt.Sprintf(" ORDER BY %s %s ", q.OrderBy, q.OrderDirection),
		1,
	)

	if m.SearchBy != "" {
		query = strings.ReplaceAll(query, "1=1", fmt.Sprintf("%v='%v'", m.SearchBy, m.Search))
	}

	args := map[string]interface{}{
		"offset": q.Offset,
		"limit":  q.Limit,
	}

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, r.translateError(err)
	}

	err = stmt.SelectContext(ctx, &result, args)
	if err != nil {
		return nil, r.translateError(err)
	}

	total, err := r.Count(ctx)
	if err != nil {
		return nil, r.translateError(err)
	}

	m.TotalItems = total

	return result, nil
}

func (r *repo) Count(ctx context.Context) (int, error) {

	result := 0

	query := `SELECT count(*) FROM cars where is_active=true`

	args := map[string]interface{}{}

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return 0, r.translateError(err)
	}

	err = stmt.GetContext(ctx, &result, args)
	if err != nil {
		return 0, r.translateError(err)
	}

	return result, nil
}

func (r *repo) CheckAvailableCars(ctx context.Context, m *meta.Params, startDate, endDate time.Time) ([]presentations.Cars, error) {

	var (
		result []presentations.Cars
	)
	query := `select 
				* 
			from 
				cars 
			where car_id not in (
				select 
					orders.car_id
				from orders where pickup_date >= :pickup_date and dropoff_date <= :dropoff_date
			) and is_active=true`

	args := map[string]any{
		"pickup_date":  startDate,
		"dropoff_date": endDate,
	}

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, r.translateError(err)
	}

	err = stmt.SelectContext(ctx, &result, args)
	if err != nil {
		return nil, r.translateError(err)
	}
	return result, nil
}
