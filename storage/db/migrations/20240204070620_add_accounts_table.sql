-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "accounts" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "account_name" varchar,
  "account_number" varchar,
  "currency" varchar,
  "status" varchar,
  "wallet_id" bigint NOT NULL,
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz DEFAULT 'now()',
  "deleted_at" timestamptz DEFAULT null
);
ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "accounts" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "accounts";
-- +goose StatementEnd
