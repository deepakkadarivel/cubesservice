-- +goose Up
CREATE TABLE IF NOT EXISTS account (
  user_id       SERIAL PRIMARY KEY,
  name          VARCHAR(255) NOT NULL,
  email         VARCHAR(255) NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  created_on    TIMESTAMP WITHOUT TIME ZONE
);

-- +goose Down
DROP TABLE IF EXISTS account CASCADE;