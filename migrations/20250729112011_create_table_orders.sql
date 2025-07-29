-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders (
    order_id            serial      NOT NULL PRIMARY KEY,
    car_id              int         NOT NULL REFERENCES cars(car_id) ON DELETE RESTRICT ON UPDATE CASCADE,
    order_date          timestamptz NOT NULL,
    pickup_date         timestamptz NOT NULL,
    dropoff_date        timestamptz NOT NULL,
    pickup_location     varchar     NOT NULL,
    dropoff_location    varchar     NOT NULL,
    is_active           boolean     NOT NULL default true,
    created_at          timestamptz NOT NULL default CURRENT_TIMESTAMP,
    updated_at          timestamptz NOT NULL default CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
