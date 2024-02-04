-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "liens" (
    "id" bigserial PRIMARY KEY,
    "wallet_id" bigint NOT NULL,
    "transaction_id" bigint NOT NULL,
    "lien_amount" bigint NOT NULL,
    "currency" varchar(50) NOT NULL,
    "status" varchar(100) NOT NULL,
    "created_at" timestamptz DEFAULT 'now()',
    "updated_at" timestamptz DEFAULT 'now()',
    "deleted_at" timestamptz DEFAULT null
    );

ALTER TABLE "liens" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "liens" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "liens";
-- +goose StatementEnd
