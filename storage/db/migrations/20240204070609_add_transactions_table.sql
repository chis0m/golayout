-- +goose Up
-- SQL to create the transactions table.
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "transactions" (
    "id" bigserial PRIMARY KEY,
    "mode" varchar(100) NOT NULL,
    "reference" varchar(100) NOT NULL,
    "transaction_type" varchar(100) NOT NULL,
    "transaction_data" jsonb NOT NULL,
    "amount" bigint NOT NULL,
    "fee" bigint NOT NULL,
    "status" varchar(100) NOT NULL,
    "source_id" int NOT NULL,
    "source_ref" varchar NOT NULL,
    "narration" varchar(255) DEFAULT null,
    "created_at" timestamptz DEFAULT 'now()',
    "updated_at" timestamptz DEFAULT 'now()',
    "deleted_at" timestamptz DEFAULT null
    );
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_transactions_updated_at_before_update
    BEFORE UPDATE ON transactions
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- SQL to remove the transactions table and its trigger.
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_transactions_updated_at_before_update ON transactions;
DROP TABLE IF EXISTS "transactions";
-- +goose StatementEnd
