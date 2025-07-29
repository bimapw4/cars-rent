-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    user_id    serial       NOT NULL PRIMARY KEY,
    username   varchar      NOT NULL UNIQUE,
    password   varchar      NOT NULL,
    is_admin   boolean      NOT NULL DEFAULT false,
    is_active  boolean      NOT NULL DEFAULT true,
    created_at timestamptz  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz  NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
