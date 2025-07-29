-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cars (
    car_id      serial      NOT NULL PRIMARY KEY,
    car_name    varchar     NOT NULL,
    day_rate    float       NOT NULL,
    month_rate  float       NOT NULL,
    image       varchar     NOT NULL,
    is_active   boolean     NOT NULL default true,
    created_at  timestamptz NOT NULL default CURRENT_TIMESTAMP,
    updated_at  timestamptz NOT NULL default CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cars;
-- +goose StatementEnd
