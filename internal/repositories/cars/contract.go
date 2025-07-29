package cars

import (
	"car-rent/internal/presentations"
	"car-rent/pkg/meta"
	"context"
)

type Cars interface {
	Create(ctx context.Context, input presentations.Cars) error
	List(ctx context.Context, m *meta.Params) ([]presentations.Cars, error)
	Detail(ctx context.Context, carsID int) (*presentations.Cars, error)
	Update(ctx context.Context, payload *presentations.Cars) error
	UpdateIsActive(ctx context.Context, carsID int, isActive bool) error
	Delete(ctx context.Context, carsID int) error
	Latest(ctx context.Context) (*presentations.Cars, error)
}
