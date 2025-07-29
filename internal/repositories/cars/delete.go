package cars

import "context"

func (r *repo) Delete(ctx context.Context, carsID int) error {

	query := `delete from cars where car_id=:car_id`

	args := map[string]interface{}{
		"car_id": carsID,
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
