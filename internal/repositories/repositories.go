package repositories

import (
	"car-rent/internal/repositories/cars"
	"car-rent/internal/repositories/order"
	"car-rent/internal/repositories/users"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Cars  cars.Cars
	Order order.Orders
	Users users.Users
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		Cars:  cars.NewRepo(db),
		Order: order.NewRepo(db),
		Users: users.NewRepo(db),
	}
}
