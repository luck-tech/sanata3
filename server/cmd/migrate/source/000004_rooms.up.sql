CREATE TABLE IF NOT EXISTS "rooms" (
  "id" varchar PRIMARY KEY,
  "name" varchar NOT NULL,
  "owner_id" varchar NOT NULL,
  "description" text,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE IF NOT EXISTS "room_members" (
  "room_id" varchar NOT NULL,
  "user_id" varchar NOT NULL,
  "joined_at" timestamp
);

CREATE TABLE IF NOT EXISTS "aim_skills" (
  "room_id" varchar NOT NULL,
  "skill_id" int NOT NULL
);

ALTER TABLE "aim_skills" ADD FOREIGN KEY ("skill_id") REFERENCES "skills" ("id");

ALTER TABLE "aim_skills" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

ALTER TABLE "rooms" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "room_members" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

ALTER TABLE "room_members" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
