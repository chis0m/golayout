CREATE USER "app-user" WITH PASSWORD 'app-password';

GRANT CONNECT ON DATABASE "app-database" TO "app-user";

GRANT USAGE ON SCHEMA public TO "app-user";

GRANT SELECT, INSERT, UPDATE ON ALL TABLES IN SCHEMA public TO "app-user";

ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE ON TABLES TO "app-user";
