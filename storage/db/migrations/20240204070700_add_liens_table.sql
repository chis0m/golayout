-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "liens" (
   "id" bigserial PRIMARY KEY,
   "wallet_id" bigint NOT NULL REFERENCES "wallets" ("id"),
    "transaction_id" bigint NOT NULL REFERENCES "transactions" ("id"),
    "lien_amount" bigint NOT NULL,
    "currency" varchar(50) NOT NULL,
    "status" varchar(100) NOT NULL,
    "created_at" timestamptz DEFAULT 'now()',
    "updated_at" timestamptz DEFAULT 'now()',
    "deleted_at" timestamptz DEFAULT null
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "liens";
-- +goose StatementEnd
