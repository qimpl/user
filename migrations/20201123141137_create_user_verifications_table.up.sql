CREATE TABLE "user_verifications" (
  "id" uuid PRIMARY KEY UNIQUE DEFAULT gen_random_uuid (),
  "stripe_verification_intent_id" varchar(50),
  "user_id" uuid REFERENCES "users" (id) ON DELETE CASCADE NOT NULL,
  "is_verified" bool DEFAULT FALSE NOT NULL,
  "status" varchar(30),
  "stripe_person_id" varchar(25),
  "verification_type" varchar(30),
  "verified_at" timestamp,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_timestamp
  BEFORE UPDATE ON "user_verifications"
  FOR EACH ROW
  EXECUTE PROCEDURE trigger_update_timestamp ();
