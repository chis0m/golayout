-- +goose Up
-- +goose StatementBegin
CREATE TABLE "post_categories" (
   "post_id" bigint NOT NULL REFERENCES "posts" ("id"),
   "category_id" bigint NOT NULL REFERENCES "categories" ("id"),
   PRIMARY KEY (post_id, category_id)
);

-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_post_categories_updated_at_before_update
    BEFORE UPDATE ON post_categories
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_post_categories_updated_at_before_update ON post_categories;
DROP TABLE IF EXISTS "categories";
-- +goose StatementEnd
