package entity

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type RoomMember struct {
	bun.BaseModel `bun:"table:room_members"`
	RoomID        string    `bun:"room_id,notnull"`
	Room          *Room     `bun:"rel:belongs-to,join:room_id=id"`
	UserID        string    `bun:"user_id,notnull"`
	User          *User     `bun:"rel:belongs-to,join:user_id=id"`
	JoinedAt      time.Time `bun:"joined_at"`
}

var _ bun.BeforeAppendModelHook = (*Room)(nil)

func (m *RoomMember) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.JoinedAt = time.Now()
	}
	return nil
}
