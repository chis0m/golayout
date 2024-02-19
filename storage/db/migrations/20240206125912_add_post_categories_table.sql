-- +goose Up
-- +goose StatementBegin
CREATE TABLE "post_categories" (
   "post_id" bigint NOT NULL REFERENCES "posts" ("id"),
   "category_id" bigint NOT NULL REFERENCES "categories" ("id"),
   PRIMARY KEY (post_id, category_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "categories";
-- +goose StatementEnd
