package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
)

type Chat struct {
	repo dai.DataAccessInterface
}

func NewChat(repo dai.DataAccessInterface) *Chat {
	return &Chat{
		repo: repo,
	}
}
func (s *Chat) Get(ctx context.Context, roomID string, lastTime time.Time) ([]entity.Chat, error) {
	return s.repo.GetChats(ctx, roomID, lastTime)
}

func (s *Chat) GetbyID(ctx context.Context, chatID string) (*entity.Chat, error) {
	return s.repo.GetChatByID(ctx, chatID)
}

func (s *Chat) Post(ctx context.Context, roomID, userID, message string) error {
	return s.repo.CreateChat(ctx, &entity.Chat{
		ID:        uuid.NewString(),
		RoomID:    roomID,
		AutherID:  userID,
		Message:   message,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	})
}

func (s *Chat) Update(ctx context.Context, chat *entity.Chat) error {
	return s.repo.UpdateChat(ctx, chat)
}
