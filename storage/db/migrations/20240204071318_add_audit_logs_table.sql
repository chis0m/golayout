-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "audit_logs" (
    "id" bigserial PRIMARY KEY,
    "user_id" bigInt NOT NULL REFERENCES "users" ("id"),
    "actor_id" Integer,
    "actor_type" Varchar,
    "action" Varchar NOT NULL,
    "description" Text,
    "status" Varchar,
    "reason" Text,
    "ip_address" Varchar,
    "resource_type" Varchar,
    "resource_id" Varchar,
    "data" Jsonb,
    "created_at" Timestamptz NOT NULL DEFAULT 'now()',
    "updated_at" Timestamptz DEFAULT 'now()',
    "deleted_at" Timestamptz DEFAULT null
    );

-- Indexes for optimized query performance
CREATE INDEX ON "audit_logs" ("user_id", "created_at");
CREATE INDEX ON "audit_logs" ("actor_id", "actor_type");
CREATE INDEX ON "audit_logs" ("action", "status");
CREATE INDEX ON "audit_logs" ("resource_type", "resource_id");
CREATE INDEX ON "audit_logs" ("ip_address");
CREATE INDEX ON "audit_logs" ("created_at");

-- Column descriptions for clarity
COMMENT ON COLUMN "audit_logs"."actor_id" IS 'ID of the actor performing the action, if different from user_id';
COMMENT ON COLUMN "audit_logs"."actor_type" IS 'Type of actor (e.g., User, System) to distinguish between users and automated processes';
COMMENT ON COLUMN "audit_logs"."action" IS 'General type of action performed (e.g., CREATE, UPDATE, DELETE)';
COMMENT ON COLUMN "audit_logs"."description" IS 'Detailed description of the action for readability';
COMMENT ON COLUMN "audit_logs"."status" IS 'Outcome of the action (SUCCESS, FAILURE)';
COMMENT ON COLUMN "audit_logs"."reason" IS 'Explanation, particularly for failures or exceptions';
COMMENT ON COLUMN "audit_logs"."ip_address" IS 'IP address from which the action was performed';
COMMENT ON COLUMN "audit_logs"."resource_type" IS 'Type of the resource being acted upon (e.g., Account, Document)';
COMMENT ON COLUMN "audit_logs"."resource_id" IS 'Identifier of the specific resource instance';
COMMENT ON COLUMN "audit_logs"."data" IS 'Additional data related to the action in JSON format';
COMMENT ON COLUMN "audit_logs"."created_at" IS 'Timestamp when the action was logged';
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "audit_logs";
-- +goose StatementEnd

