package repository

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
	"github.com/uptrace/bun"
)

type SessionRepository struct {
	db bun.IDB
}

func NewSessionRepository(db bun.IDB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (s *SessionRepository) CreateSeseion(ctx context.Context, session *entity.Session) error {
	panic("impl me")
}

func (s *SessionRepository) GetSessionByID(ctx context.Context, sessionID string) (*entity.Session, bool, error) {
	panic("impl me")
}

func (s *SessionRepository) GetSessionByUserID(ctx context.Context, userID string) (*entity.Session, bool, error) {
	panic("impl me")
}

func (s *SessionRepository) UpdateSession(ctx context.Context, session *entity.Session) error {
	panic("impl me")
}

func (s *SessionRepository) DeleteSession(ctx context.Context, id string) error {
	panic("impl me")
}

var _ dai.Session = (*SessionRepository)(nil)
