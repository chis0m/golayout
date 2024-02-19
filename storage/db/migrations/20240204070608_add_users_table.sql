-- +goose Up
-- SQL to create the users table and its trigger for automatic updated_at handling.
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "users" (
    "id" bigserial PRIMARY KEY,
    "first_name" varchar,
    "last_name" varchar,
    "email" varchar UNIQUE NOT NULL,
    "email_verified_at" timestamptz DEFAULT NULL,
    "password_hash" varchar NOT NULL,
    "address" varchar,
    "bvn" varchar,
    "created_at" timestamptz DEFAULT 'now()',
    "updated_at" timestamptz DEFAULT 'now()',
    "deleted_at" timestamptz DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
