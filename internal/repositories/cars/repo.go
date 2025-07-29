package cars

import (
	"car-rent/internal/presentations"
	"car-rent/pkg/databasex"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) Cars {
	return &repo{
		db: db,
	}
}

func (r *repo) translateError(err error) error {
	switch err {
	case sql.ErrNoRows:
		return presentations.ErrCarsNotExist
	case databasex.ErrUniqueViolation:
		return presentations.ErrCarsAlreadyExist
	default:
		return err
	}
}
