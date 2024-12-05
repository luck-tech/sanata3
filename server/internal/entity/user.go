package entity

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID          string     `bun:"id,pk"`
	Name        string     `bun:"name,notnull"`
	Email       string     `bun:"email,notnull"`
	Icon        string     `bun:"icon,notnull"`
	Description string     `bun:"description"`
	CreatedAt   time.Time  `bun:"created_at"`
	UpdatedAt   time.Time  `bun:"updated_at"`
	DeletedAt   *time.Time `bun:"deleted_at,soft_delete"`
}

var _ bun.BeforeAppendModelHook = (*User)(nil)

func (m *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

type UserSkills struct {
	User           *User
	UsedSkill      []UsedSkill
	WantLeanSkills []WantLearnSkill
}
