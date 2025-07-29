-- +goose Up
-- +goose StatementBegin
INSERT INTO users (username, password, is_admin, created_at, updated_at) VALUES
  ('admin_user', '$2a$10$01CFJTdZbxF7UjD45C31Ne2O4Gt11CmIBBuA8mREIpbcUrDT58TfK', true, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('user_one',   '$2a$10$01CFJTdZbxF7UjD45C31Ne2O4Gt11CmIBBuA8mREIpbcUrDT58TfK', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('user_two',   '$2a$10$01CFJTdZbxF7UjD45C31Ne2O4Gt11CmIBBuA8mREIpbcUrDT58TfK', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('user_three', '$2a$10$01CFJTdZbxF7UjD45C31Ne2O4Gt11CmIBBuA8mREIpbcUrDT58TfK', false, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
