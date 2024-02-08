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

-- +goose StatementBegin
CREATE TRIGGER update_comments_updated_at_before_update
    BEFORE UPDATE ON comments
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_comments_updated_at_before_update ON comments;
DROP TABLE IF EXISTS "comments";
-- +goose StatementEnd
