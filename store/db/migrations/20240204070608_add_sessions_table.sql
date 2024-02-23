-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "sessions" (
    "id" uuid PRIMARY KEY,
    "user_id" bigint NOT NULL REFERENCES "users" ("id"),
    "refresh_token" varchar NOT NULL,
    "user_agent" varchar NOT NULL,
    "client_ip" varchar NOT NULL,
    "is_blocked" boolean NOT NULL DEFAULT false,
    "expires_at" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "deleted_at" timestamptz DEFAULT (NULL)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "sessions";
-- +goose StatementEnd
