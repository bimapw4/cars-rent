package order

import "context"

func (r *repo) Delete(ctx context.Context, orderID int) error {

	query := `delete from orders where order_id=:order_id`

	args := map[string]interface{}{
		"order_id": orderID,
	}
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return r.translateError(err)
	}

	_, err = stmt.ExecContext(ctx, args)
	if err != nil {
		return r.translateError(err)
	}

	return nil
}
