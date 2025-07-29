package handlers

import (
	"car-rent/internal/business"
	"car-rent/internal/handlers/cars"
)

type Handlers struct {
	Cars cars.Handler
}

func NewHandler(business business.Business) Handlers {
	return Handlers{
		Cars: cars.NewHandler(business),
	}
}
