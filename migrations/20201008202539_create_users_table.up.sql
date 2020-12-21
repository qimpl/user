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
  "civility" varchar(4) DEFAULT 'Mr' NOT NULL,
  "first_name" varchar(64) NOT NULL,
  "last_name" varchar(64) NOT NULL,
  "email" varchar(100) NOT NULL UNIQUE,
  "mobile_phone_number" varchar(20) NOT NULL UNIQUE,
  "password" varchar(100) NOT NULL,
  "birthdate" timestamp NOT NULL,
  "country" varchar(20) NOT NULL,
  "state" varchar(50),
  "street" varchar(100) NOT NULL,
  "additional_street_information" varchar(100),
  "city" varchar(64) NOT NULL,
  "zip" varchar(20) NOT NULL,
  "is_owner" bool DEFAULT FALSE NOT NULL,
  "is_enabled" bool DEFAULT FALSE NOT NULL,
  "is_admin" bool DEFAULT FALSE NOT NULL,
  "is_deleted" bool DEFAULT FALSE NOT NULL,
  "stripe_customer_id" varchar(20),
  "stripe_account_id" varchar(25),
  "stripe_payment_method_id" varchar(30),
  "push_notification_token" varchar(50),
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp
);

CREATE TRIGGER update_timestamp
  BEFORE UPDATE ON "users"
  FOR EACH ROW
  EXECUTE PROCEDURE trigger_update_timestamp ();
