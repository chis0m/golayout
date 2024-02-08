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

-- +goose StatementBegin
CREATE TRIGGER update_users_updated_at_before_update
    BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- SQL to remove the users table and its trigger.
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_users_updated_at_before_update ON users;
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
