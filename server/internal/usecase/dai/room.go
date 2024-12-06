package dai

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
)

type Room interface {
	CreateRoom(ctx context.Context, room *entity.Room) error
	UpdateRoom(ctx context.Context, room *entity.Room) error
	GetRoom(ctx context.Context, roomID string) (*entity.Room, bool, error)
	ListRoom(ctx context.Context, userID string) ([]entity.Room, error)
	DeleteRoom(ctx context.Context, roomID string) error
}
