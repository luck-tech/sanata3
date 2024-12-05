CREATE TABLE IF NOT EXISTS "skills" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS "want_learn_skills" (
  "user_id" varchar NOT NULL,
  "skill_id" int NOT NULL
);

CREATE TABLE IF NOT EXISTS "used_skills" (
  "user_id" varchar NOT NULL,
  "skill_id" int NOT NULL
);

ALTER TABLE "want_learn_skills" ADD FOREIGN KEY ("skill_id") REFERENCES "skills" ("id");

ALTER TABLE "used_skills" ADD FOREIGN KEY ("skill_id") REFERENCES "skills" ("id");

ALTER TABLE "want_learn_skills" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "used_skills" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
