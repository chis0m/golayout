-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "users" (
 "id" bigserial PRIMARY KEY,
 "first_name" varchar,
 "last_name" varchar,
 "email" varchar UNIQUE NOT NULL,
 "email_verified_at" timestamptz,
 "address" varchar,
 "bvn" varchar,
 "created_at" timestamptz DEFAULT 'now()',
 "updated_at" timestamptz DEFAULT 'now()',
 "deleted_at" timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users"
-- +goose StatementEnd
