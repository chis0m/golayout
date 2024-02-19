-- +goose Up
-- SQL to create the accounts table and setup foreign keys.
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "accounts" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL REFERENCES "users" ("id"),
  "wallet_id" bigint NOT NULL REFERENCES "wallets" ("id"),
  "account_name" varchar,
  "account_number" varchar,
  "currency" varchar,
  "status" varchar,
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz DEFAULT 'now()',
  "deleted_at" timestamptz DEFAULT null
);
-- +goose StatementEnd

-- +goose Down
-- SQL to remove the accounts table trigger and the table itself.
-- +goose StatementBegin
DROP TABLE IF EXISTS "accounts";
-- +goose StatementEnd
