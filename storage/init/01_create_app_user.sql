CREATE DATABASE "app-database";

CREATE USER "app-user" WITH PASSWORD 'app-password';

GRANT CONNECT ON DATABASE "app-database" TO "app-user";

-- This should be executed within the "app-database" context
\c "app-database";

GRANT USAGE, CREATE ON SCHEMA public TO "app-user";

GRANT SELECT, INSERT, UPDATE ON ALL TABLES IN SCHEMA public TO "app-user";
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, REFERENCES ON TABLES TO "app-user";

GRANT USAGE, UPDATE ON ALL SEQUENCES IN SCHEMA public TO "app-user";
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT USAGE, UPDATE ON SEQUENCES TO "app-user";
