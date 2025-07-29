package users

import (
	"car-rent/internal/presentations"
	"context"
	"time"
)

func (r *repo) UpdatePassword(ctx context.Context, userID, password, updatedBy string) error {

	query := `
	update users set password=:password, updated_at=:updated_at, updated_by=:updated_by where id=:id
   	`

	// Execute the query using named parameters
	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":         userID,
		"password":   password,
		"updated_by": updatedBy,
		"updated_at": time.Now(),
	})

	return err
}

func (r *repo) Update(ctx context.Context, payload presentations.Users) error {

	query := `
		update users set 
			username=:username, 
			password=:password, 
			id_admin=:id_admin, 
			updated_at=:updated_at 
		where id=:id
   	`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":         payload.UserID,
		"username":   payload.Username,
		"password":   payload.Password,
		"is_admin":   payload.IsAdmin,
		"updated_at": time.Now(),
	})

	return err
}

func (r *repo) DeleteUser(ctx context.Context, userID, updatedBy string) error {

	query := `
		update users set 
			is_active=:is_active, 
			updated_at=:updated_at, 
			updated_by=:updated_by 
		where id=:id
   	`

	// Execute the query using named parameters
	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":         userID,
		"is_active":  false,
		"updated_by": updatedBy,
		"updated_at": time.Now(),
	})

	return err
}
