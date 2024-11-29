package dai

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
)

type Session interface {
	CreateSeseion(ctx context.Context, session *entity.Session) error
	GetSessionByID(ctx context.Context, sessionID string) (*entity.Session, bool, error)
	GetSessionByUserID(ctx context.Context, userID string) (*entity.Session, bool, error)
	UpdateSession(ctx context.Context, session *entity.Session) error
	DeleteSession(ctx context.Context, id string) error
}
