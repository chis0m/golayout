-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "wallet_history" (
    "id" int PRIMARY KEY,
    "wallet_id" int NOT NULL REFERENCES "wallets" ("id"),
    "public_id" varchar NOT NULL,
    "balance" int NOT NULL,
    "currency" varchar(50) NOT NULL,
    "change_amount" int NOT NULL,
    "locked_amount" int NOT NULL,
    "mode" varchar(100) NOT NULL,
    "hash" varchar(255) NOT NULL,
    "operation" varchar(100) NOT NULL,
    "previous_hash" varchar(255) DEFAULT null,
    "transaction_id" integer NOT NULL REFERENCES "transactions" ("id"),
    "created_at" timestamptz DEFAULT 'now()',
    "updated_at" timestamptz DEFAULT 'now()',
    "deleted_at" timestamptz DEFAULT null
    );
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_wallet_history_updated_at_before_update
    BEFORE UPDATE ON "wallet_history"
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_wallet_history_updated_at_before_update ON "wallet_history";
DROP TABLE IF EXISTS "wallet_history";
-- +goose StatementEnd
