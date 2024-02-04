-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "wallets" (
    "id" bigserial PRIMARY KEY,
    "uuid" varchar NOT NULL,
    "balance" bigint NOT NULL,
    "currency" varchar(50) NOT NULL,
    "change_amount" bigint NOT NULL,
    "locked_amount" bigint NOT NULL,
    "status" varchar(100) NOT NULL,
    "mode" varchar(100) NOT NULL,
    "hash" varchar(255) NOT NULL,
    "previous_hash" varchar(255) DEFAULT null,
    "transaction_id" bigint NOT NULL,
    "created_at" timestamptz DEFAULT 'now()',
    "updated_at" timestamptz DEFAULT 'now()',
    "deleted_at" timestamptz DEFAULT null
    );

ALTER TABLE "wallets" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "wallets";
-- +goose StatementEnd
