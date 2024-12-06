package dai

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
)

type RoomMember interface {
	GetRoomMembers(ctx context.Context, roomID string) ([]entity.RoomMember, error)
	GetRoomMembersByRoomIDs(ctx context.Context, roomIDs []string) ([]entity.RoomMember, error)
	JoinRoom(ctx context.Context, roomID, userID string) error
	LeaveRoom(ctx context.Context, roomID, userID string) error
	DeleteRoomMembers(ctx context.Context,roomID string)error
}
