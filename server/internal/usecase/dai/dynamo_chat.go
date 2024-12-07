package dai

import (
	"context"
	"time"

	"github.com/murasame29/go-httpserver-template/internal/entity"
)

type Chat interface {
	CreateChat(ctx context.Context, chat *entity.Chat) error
	GetChats(ctx context.Context, roomID string, lastTime time.Time) ([]entity.Chat, error)
	GetChatByID(ctx context.Context, chatID string) (*entity.Chat, error)
	UpdateChat(ctx context.Context, chat *entity.Chat) error
}
