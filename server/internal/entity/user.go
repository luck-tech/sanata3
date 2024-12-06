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

type UserInfo struct {
	User           *User
	UsedSkill      []UsedSkill
	WantLeanSkills []WantLearnSkill
}

func ToUserMap(users []User) map[string]User {
	userMap := make(map[string]User)
	for _, user := range users {
		userMap[user.ID] = user
	}
	return userMap
}

type DisplayUser struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
}

func ToDisplayUser(user User) DisplayUser {
	return DisplayUser{
		ID:          user.ID,
		Name:        user.Name,
		Icon:        user.Icon,
		Description: user.Description,
	}
}

func ToDisplayUsers(users []User) []DisplayUser {
	result := make([]DisplayUser, len(users))
	for i, user := range users {
		result[i] = ToDisplayUser(user)
	}
	return result
}
