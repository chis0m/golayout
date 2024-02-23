-- +goose Up
-- +goose StatementBegin
CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" VARCHAR(255) UNIQUE NOT NULL,
  "deleted_at" timestamptz DEFAULT null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "categories";
-- +goose StatementEnd
