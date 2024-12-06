package entity

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Room struct {
	bun.BaseModel `bun:"table:rooms"`
	ID            string     `bun:"id,pk"`
	Name          string     `bun:"name,notnull"`
	OwnerID       string     `bun:"owner_id,notnull"`
	Owner         *User      `bun:"rel:belongs-to,join:owner_id=id"`
	Description   string     `bun:"description"`
	CreatedAt     time.Time  `bun:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete"`
}

var _ bun.BeforeAppendModelHook = (*Room)(nil)

func (m *Room) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}

type RoomInfo struct {
	Room       *Room
	AimSkill   []AimSkill
	RoomMember []RoomMember
}
