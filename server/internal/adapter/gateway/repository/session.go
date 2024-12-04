package repository

import (
	"context"
	"database/sql"
	"errors"

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
	if _, err := s.db.NewInsert().Model(session).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *SessionRepository) GetSessionByID(ctx context.Context, sessionID string) (*entity.Session, bool, error) {
	var session entity.Session
	query := s.db.NewSelect().Model(&session).
		Where("id = ?", sessionID)
		// Where("user_agent = ?", userAgent)

	if err := query.Scan(ctx, &session); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, false, err
		}
		return nil, false, nil
	}

	return &session, true, nil
}

func (s *SessionRepository) GetSessionByUserID(ctx context.Context, userID string) (*entity.Session, bool, error) {
	var session entity.Session
	query := s.db.NewSelect().Model(&session).
		Where("user_id = ?", userID)
		// Where("user_agent = ?", userAgent)

	if err := query.Scan(ctx, &session); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, false, err
		}
		return nil, false, nil
	}

	return &session, true, nil
}

func (s *SessionRepository) UpdateSession(ctx context.Context, session *entity.Session) error {
	if _, err := s.db.NewUpdate().Model(session).WherePK().Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *SessionRepository) DeleteSession(ctx context.Context, id string) error {
	if _, err := s.db.NewDelete().Model(&entity.User{}).Where("id = ?", id).Exec(ctx); err != nil {
		return err
	}

	return nil
}

var _ dai.Session = (*SessionRepository)(nil)
