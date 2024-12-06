package service

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
)

type RoomMember struct {
	repo dai.DataAccessInterface
}

func NewRoomMember(
	repo dai.DataAccessInterface,
) *RoomMember {
	return &RoomMember{
		repo: repo,
	}
}

func (s *RoomMember) List(ctx context.Context, roomID string) ([]entity.RoomMember, error) {
	return s.repo.GetRoomMembers(ctx, roomID)
}

type ListByRoomIDsResult struct {
	Members map[string][]entity.RoomMember
	UserIDs []string
}

func (s *RoomMember) ListByRoomIDs(ctx context.Context, roomIDs []string) (*ListByRoomIDsResult, error) {
	roomMembers, err := s.repo.GetRoomMembersByRoomIDs(ctx, roomIDs)
	if err != nil {
		return nil, err
	}

	result := make(map[string][]entity.RoomMember)
	userIDsMap := make(map[string]struct{})
	for _, roomMember := range roomMembers {
		result[roomMember.RoomID] = append(result[roomMember.RoomID], roomMember)
		userIDsMap[roomMember.UserID] = struct{}{}
	}

	var userIDs []string
	for skillID := range userIDsMap {
		userIDs = append(userIDs, skillID)
	}

	return &ListByRoomIDsResult{
		Members: result,
		UserIDs: userIDs,
	}, nil
}

func (s *RoomMember) Join(ctx context.Context, roomID, userID string) error {
	return s.repo.JoinRoom(ctx, roomID, userID)
}

func (s *RoomMember) Leave(ctx context.Context, roomID, userID string) error {
	return s.repo.LeaveRoom(ctx, roomID, userID)
}
