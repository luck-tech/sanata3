package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
)

type Session struct {
	repo dai.DataAccessInterface
}

func NewSession(repo dai.DataAccessInterface) *Session {
	return &Session{
		repo: repo,
	}
}

func (s *Session) GetSessionByID(ctx context.Context, sessionID string) (*entity.Session, bool, error) {
	return s.repo.GetSessionByID(ctx, sessionID)
}

func (s *Session) UpsertSession(ctx context.Context, userID, accessToken, refreshToken string) (string, error) {
	session, found, err := s.repo.GetSessionByUserID(ctx, userID)
	if err != nil {
		return "", err
	}

	if !found {
		newSession := &entity.Session{
			ID:           uuid.NewString(),
			UserID:       userID,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}

		if err := s.repo.CreateSeseion(ctx, newSession); err != nil {
			return "", err
		}
	} else {
		session.AccessToken = accessToken
		session.RefreshToken = refreshToken

		if err := s.repo.UpdateSession(ctx, session); err != nil {
			return "", err
		}
	}

	return session.ID, nil
}
