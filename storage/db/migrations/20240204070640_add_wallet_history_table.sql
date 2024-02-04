-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "wallet_history" (
    "id" int PRIMARY KEY,
    "wallet_id" int NOT NULL,
    "public_id" varchar NOT NULL,
    "balance" int NOT NULL,
    "currency" varchar(50) NOT NULL,
    "change_amount" int NOT NULL,
    "locked_amount" int NOT NULL,
    "mode" varchar(100) NOT NULL,
    "hash" varchar(255) NOT NULL,
    "operation" varchar(100) NOT NULL,
    "previous_hash" varchar(255) DEFAULT null,
    "transaction_id" integer NOT NULL,
    "created_at" timestamptz DEFAULT 'now()',
    "updated_at" timestamptz DEFAULT 'now()',
    "deleted_at" timestamptz DEFAULT null
    );

ALTER TABLE "wallet_history" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "wallet_history" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "wallet_history";
-- +goose StatementEnd
