package business

import (
	"car-rent/internal/business/auth"
	"car-rent/internal/business/cars"
	order "car-rent/internal/business/orders"
	"car-rent/internal/repositories"
)

type Business struct {
	Cars  cars.Cars
	Order order.Orders
	Auth  auth.Contract
}

func NewBusiness(repo *repositories.Repository) Business {
	return Business{
		Cars:  cars.NewBusiness(repo),
		Order: order.NewBusiness(repo),
		Auth:  auth.NewBusiness(repo),
	}
}
