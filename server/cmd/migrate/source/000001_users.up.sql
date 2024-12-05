CREATE TABLE IF NOT EXISTS "users" (
  "id" varchar PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "icon" varchar NOT NULL,
  "description" text,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);