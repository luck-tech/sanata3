package entity

import "github.com/uptrace/bun"

type WantLearnSkill struct {
	bun.BaseModel `bun:"table:want_learn_skills,alias:wls"`
	UserID        string `bun:"user_id,notnull"`
	User          *User  `bun:"rel:belongs-to,join:user_id=id"`
	SkillID       int    `bun:"skill_id,notnull"`
	Skill         *Skill `bun:"rel:belongs-to,join:skill_id=id"`
}
