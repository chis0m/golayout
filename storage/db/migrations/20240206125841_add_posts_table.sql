-- +goose Up
-- +goose StatementBegin
CREATE TABLE "posts" (
 "id" bigserial PRIMARY KEY,
 "user_id" bigint NOT NULL REFERENCES "users" ("id"),
 "title" VARCHAR NOT NULL,
 "content" TEXT NOT NULL,
 "created_at" timestamptz DEFAULT (now()),
 "updated_at" timestamptz DEFAULT (now()),
 "deleted_at" timestamptz DEFAULT null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "posts";
-- +goose StatementEnd
