-- +goose Up
-- +goose StatementBegin
CREATE TABLE "comments" (
"id" bigserial PRIMARY KEY,
"post_id" bigint NOT NULL REFERENCES "posts" ("id"),
"user_id" bigint NOT NULL REFERENCES "users" ("id"),
"content" TEXT NOT NULL,
"created_at" timestamptz DEFAULT (now()),
"deleted_at" timestamptz DEFAULT null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "comments";
-- +goose StatementEnd
