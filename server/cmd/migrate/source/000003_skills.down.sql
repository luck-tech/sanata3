ALTER TABLE "want_learn_skills" DROP CONSTRAINT IF EXISTS "want_learn_skills_skill_id_fkey";

ALTER TABLE "used_skills" DROP CONSTRAINT IF EXISTS "used_skills_skill_id_fkey";

ALTER TABLE "want_learn_skills" DROP CONSTRAINT IF EXISTS "want_learn_skills_user_id_fkey";

ALTER TABLE "used_skills" DROP CONSTRAINT IF EXISTS "used_skills_user_id_fkey";

DROP TABLE IF EXISTS "want_learn_skills";

DROP TABLE IF EXISTS "used_skills";

DROP TABLE IF EXISTS "skills";