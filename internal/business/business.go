package business

import (
	"car-rent/internal/repositories"
)

type Business struct {
}

func NewBusiness(repo *repositories.Repository) Business {
	return Business{}
}
