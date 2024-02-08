-- +goose Up
-- +goose StatementBegin
CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" VARCHAR(255) UNIQUE NOT NULL,
  "deleted_at" timestamptz DEFAULT null
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TRIGGER update_categories_updated_at_before_update
    BEFORE UPDATE ON categories
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS update_categories_updated_at_before_update ON categories;
DROP TABLE IF EXISTS "categories";
-- +goose StatementEnd
