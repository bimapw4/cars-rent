package order

import (
	"car-rent/internal/presentations"
	"car-rent/pkg/meta"
	"context"
)

type Orders interface {
	Create(ctx context.Context, input presentations.Order) error
	List(ctx context.Context, m *meta.Params) ([]presentations.Order, error)
	Detail(ctx context.Context, orderID int) (*presentations.Order, error)
	Update(ctx context.Context, payload *presentations.Order) error
	UpdateIsActive(ctx context.Context, orderID int, isActive bool) error
	Delete(ctx context.Context, orderID int) error
	Latest(ctx context.Context) (*presentations.Order, error)
	DetailWithoutIsActive(ctx context.Context, orderID int) (*presentations.Order, error)
}
