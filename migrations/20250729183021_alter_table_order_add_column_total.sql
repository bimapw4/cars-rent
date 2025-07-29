-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders ADD COLUMN IF NOT EXISTS total_payment float NOT NULL default 0;
ALTER TABLE orders ADD COLUMN IF NOT EXISTS user_id int NOT NULL default 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE orders DROP COLUMN IF EXISTS total_payment;
ALTER TABLE orders DROP COLUMN IF EXISTS user_id;
-- +goose StatementEnd
