package handlers

import (
	"car-rent/internal/business"
	"car-rent/internal/handlers/auth"
	"car-rent/internal/handlers/cars"
	"car-rent/internal/handlers/orders"
)

type Handlers struct {
	Cars  cars.Handler
	Order orders.Handler
	Auth  auth.Handler
}

func NewHandler(business business.Business) Handlers {
	return Handlers{
		Cars:  cars.NewHandler(business),
		Order: orders.NewHandler(business),
		Auth:  auth.NewHandler(business),
	}
}
