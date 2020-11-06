CREATE TYPE weekdays AS ENUM (
  '0',
  '1',
  '2',
  '3',
  '4',
  '5',
  '6'
);

CREATE TABLE "time_slots" (
  "id" uuid PRIMARY KEY UNIQUE DEFAULT gen_random_uuid (),
  "weekday" weekdays NOT NULL,
  "start_time" time NOT NULL,
  "end_time" time NOT NULL,
  "user_id" uuid REFERENCES "users" (id) ON DELETE CASCADE NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  UNIQUE ("weekday", "start_time", "end_time", "user_id")
);

CREATE TRIGGER update_timestamp
  BEFORE UPDATE ON "time_slots"
  FOR EACH ROW
  EXECUTE PROCEDURE trigger_update_timestamp ();
