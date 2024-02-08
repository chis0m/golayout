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

-- +goose StatementBegin
CREATE TRIGGER update_liens_updated_at_before_update
    BEFORE UPDATE ON "liens"
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_liens_updated_at_before_update ON "liens";
DROP TABLE IF EXISTS "liens";
-- +goose StatementEnd
