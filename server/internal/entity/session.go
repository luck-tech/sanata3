package entity

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Session struct {
	bun.BaseModel `bun:"table:sessions"`

	ID           string    `bun:"id,pk"`
	UserID       string    `bun:"user_id,notnull"`
	User         *User     `bun:"rel:belongs-to,join:user_id=id"`
	UserAgent    string    `bun:"user_agent,notnull"`
	AccessToken  string    `bun:"access_token,notnull"`
	RefreshToken string    `bun:"refresh_token,notnull"`
	CreatedAt    time.Time `bun:"created_at"`
	UpdatedAt    time.Time `bun:"updated_at"`
}

var _ bun.BeforeAppendModelHook = (*Session)(nil)

func (m *Session) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}
