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
    "transaction_id" bigint NOT NULL REFERENCES "transactions" ("id"),
    "created_at" timestamptz DEFAULT 'now()',
    "updated_at" timestamptz DEFAULT 'now()',
    "deleted_at" timestamptz DEFAULT null
    );
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_wallets_updated_at_before_update
    BEFORE UPDATE ON "wallets"
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_wallets_updated_at_before_update ON "wallets";
DROP TABLE IF EXISTS "wallets";
-- +goose StatementEnd

