-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2024-02-04T07:32:32.706Z

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar,
  "last_name" varchar,
  "email" varchar UNIQUE NOT NULL,
  "email_verified_at" timestamptz,
  "address" varchar,
  "bvn" varchar,
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz DEFAULT 'now()',
  "deleted_at" timestamptz
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "account_name" varchar,
  "account_number" varchar,
  "currency" varchar,
  "status" varchar,
  "wallet_id" bigint NOT NULL,
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz DEFAULT 'now()',
  "deleted_at" timestamptz DEFAULT null
);

CREATE TABLE "wallets" (
  "id" bigserial PRIMARY KEY,
  "uuid" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar(50) NOT NULL,
  "change_amount" bigint NOT NULL,
  "locked_amount" bigint NOT NULL,
  "status" varchar(100) NOT NULL,
  "mode" varchar(100) NOT NULL,
  "hash" varchar(255) NOT NULL,
  "previous_hash" varchar(255) DEFAULT null,
  "transaction_id" bigint NOT NULL,
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz DEFAULT 'now()',
  "deleted_at" timestamptz DEFAULT null
);

CREATE TABLE "wallet_history" (
  "id" int PRIMARY KEY,
  "wallet_id" int NOT NULL,
  "public_id" varchar NOT NULL,
  "balance" int NOT NULL,
  "currency" varchar(50) NOT NULL,
  "change_amount" int NOT NULL,
  "locked_amount" int NOT NULL,
  "mode" varchar(100) NOT NULL,
  "hash" varchar(255) NOT NULL,
  "operation" varchar(100) NOT NULL,
  "previous_hash" varchar(255) DEFAULT null,
  "transaction_id" integer NOT NULL,
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz DEFAULT 'now()',
  "deleted_at" timestamptz DEFAULT null
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "mode" varchar(100) NOT NULL,
  "reference" varchar(100) NOT NULL,
  "transaction_type" varchar(100) NOT NULL,
  "transaction_data" jsonb NOT NULL,
  "amount" bigint NOT NULL,
  "fee" bigint NOT NULL,
  "status" varchar(100) NOT NULL,
  "source_id" int NOT NULL,
  "source_ref" varchar NOT NULL,
  "narration" varchar(255) DEFAULT null,
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz DEFAULT 'now()',
  "deleted_at" timestamptz DEFAULT null
);

CREATE TABLE "liens" (
  "id" bigserial PRIMARY KEY,
  "wallet_id" bigint NOT NULL,
  "transaction_id" bigint NOT NULL,
  "lien_amount" bigint NOT NULL,
  "currency" varchar(50) NOT NULL,
  "status" varchar(100) NOT NULL,
  "created_at" timestamptz DEFAULT 'now()',
  "updated_at" timestamptz DEFAULT 'now()',
  "deleted_at" timestamptz DEFAULT null
);

CREATE TABLE "audit_logs" (
  "id" bigSerial PRIMARY KEY,
  "user_id" BigInt NOT NULL,
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
  "created_at" Timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "audit_logs" ("user_id", "created_at");

CREATE INDEX ON "audit_logs" ("actor_id", "actor_type");

CREATE INDEX ON "audit_logs" ("action", "status");

CREATE INDEX ON "audit_logs" ("resource_type", "resource_id");

CREATE INDEX ON "audit_logs" ("ip_address");

CREATE INDEX ON "audit_logs" ("created_at");

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

ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "accounts" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "wallets" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");

ALTER TABLE "wallet_history" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "wallet_history" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");

ALTER TABLE "liens" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "liens" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");

ALTER TABLE "audit_logs" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
