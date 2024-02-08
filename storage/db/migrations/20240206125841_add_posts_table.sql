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

-- +goose StatementBegin
CREATE TRIGGER update_posts_updated_at_before_update
    BEFORE UPDATE ON posts
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_posts_updated_at_before_update ON posts;
DROP TABLE IF EXISTS "posts";
-- +goose StatementEnd
