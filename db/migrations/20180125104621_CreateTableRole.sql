-- +goose Up
CREATE TABLE IF NOT EXISTS role (
  role_id   SERIAL PRIMARY KEY,
  role_name VARCHAR(50) UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS role CASCADE;
