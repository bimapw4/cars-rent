package users

import (
	"car-rent/internal/presentations"
	"context"
)

func (r *repo) Create(ctx context.Context, input presentations.Users) error {

	query := `
    INSERT INTO users (
        username, password, is_admin, created_at, updated_at
    ) VALUES (
        :username, :password, :is_admin, :created_at, :updated_at
    )`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"username":   input.Username,
		"password":   input.Password,
		"is_admin":   input.IsAdmin,
		"created_at": input.CreatedAt,
		"updated_at": input.UpdatedAt,
	})

	return err
}
