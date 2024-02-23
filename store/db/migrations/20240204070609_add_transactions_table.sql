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

-- +goose Down
-- SQL to remove the transactions table and its trigger.
-- +goose StatementBegin
DROP TABLE IF EXISTS "transactions";
-- +goose StatementEnd
