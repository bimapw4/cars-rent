package repositories

import (
	"car-rent/internal/repositories/cars"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Cars cars.Cars
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		Cars: cars.NewRepo(db),
	}
}
