package business

import (
	"car-rent/internal/business/cars"
	"car-rent/internal/repositories"
)

type Business struct {
	Cars cars.Cars
}

func NewBusiness(repo *repositories.Repository) Business {
	return Business{
		Cars: cars.NewBusiness(repo),
	}
}
