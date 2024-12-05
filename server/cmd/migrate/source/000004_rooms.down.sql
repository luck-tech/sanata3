ALTER TABLE "aim_skills" DROP CONSTRAINT IF EXISTS "aim_skills_skill_id_fkey";

ALTER TABLE "aim_skills" DROP CONSTRAINT IF EXISTS "aim_skills_room_id_fkey";

ALTER TABLE "rooms" DROP CONSTRAINT IF EXISTS "rooms_owner_id_fkey";

ALTER TABLE "room_members" DROP CONSTRAINT IF EXISTS "room_members_room_id_fkey";

ALTER TABLE "room_members" DROP CONSTRAINT IF EXISTS "room_members_user_id_fkey";

DROP TABLE IF EXISTS "aim_skills";

DROP TABLE IF EXISTS "room_members";

DROP TABLE IF EXISTS "rooms";