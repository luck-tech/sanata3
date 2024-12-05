package entity

import "github.com/uptrace/bun"

type AimSkill struct {
	bun.BaseModel `bun:"table:aim_skills"`
	RoomID        string `bun:"room_id,notnull"`
	Room          *Room  `bun:"rel:belongs-to,join:room_id=id"`
	SkillID       int    `bun:"skill_id,notnull"`
	Skill         *Skill `bun:"rel:belongs-to,join:skill_id=id"`
}
