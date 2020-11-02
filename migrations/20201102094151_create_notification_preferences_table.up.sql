CREATE TABLE "notification_preferences" (
    "id" uuid PRIMARY KEY UNIQUE DEFAULT gen_random_uuid (),
    "user_id" uuid REFERENCES "users" (id) ON DELETE CASCADE NOT NULL,
    "on_email" bool DEFAULT TRUE NOT NULL,
    "on_sms" bool DEFAULT FALSE NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_timestamp
    BEFORE UPDATE ON "notification_preferences"
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_update_timestamp ();

