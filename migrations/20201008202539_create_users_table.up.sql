CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE OR REPLACE FUNCTION trigger_update_timestamp ()
    RETURNS TRIGGER
    AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$
LANGUAGE plpgsql;

CREATE TABLE "users" (
    "id" uuid PRIMARY KEY UNIQUE DEFAULT gen_random_uuid (),
    "first_name" varchar(64) NOT NULL,
    "last_name" varchar(64) NOT NULL,
    "email" varchar(100) NOT NULL UNIQUE,
    "password" varchar(100) NOT NULL,
    "birthdate" timestamp NOT NULL,
    "country" varchar(20) NOT NULL,
    "state" varchar(50),
    "street" varchar(100) NOT NULL,
    "city" varchar(64) NOT NULL,
    "zip" varchar(20) NOT NULL,
    "is_owner" bool DEFAULT FALSE NOT NULL,
    "is_enabled" bool DEFAULT FALSE NOT NULL,
    "is_admin" bool DEFAULT FALSE NOT NULL,
    "is_deleted" bool DEFAULT FALSE NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_timestamp
    BEFORE UPDATE ON "users"
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_update_timestamp ();

