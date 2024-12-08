package repository

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
	"github.com/uptrace/bun"
)

type RoomMemberRepository struct {
	db bun.IDB
}

func NewRoomMemberRepository(db bun.IDB) *RoomMemberRepository {
	return &RoomMemberRepository{db: db}
}

func (r *RoomMemberRepository) GetRoomMembers(ctx context.Context, roomID string) ([]entity.RoomMember, error) {
	roomMembers := []entity.RoomMember{}
	if err := r.db.NewSelect().Model(&roomMembers).Where("room_id = ?", roomID).Scan(ctx, &roomMembers); err != nil {
		return nil, err
	}

	return roomMembers, nil
}

func (r *RoomMemberRepository) GetRoomMembersByRoomIDs(ctx context.Context, roomIDs []string) ([]entity.RoomMember, error) {
	roomMembers := []entity.RoomMember{}
	if err := r.db.NewSelect().Model(&roomMembers).Where("room_id IN (?)", bun.In(roomIDs)).Scan(ctx, &roomMembers); err != nil {
		return nil, err
	}

	return roomMembers, nil
}

func (r *RoomMemberRepository) JoinRoom(ctx context.Context, roomID, userID string) error {
	newRoomMember := &entity.RoomMember{
		RoomID: roomID,
		UserID: userID,
	}

	if _, err := r.db.NewInsert().Model(newRoomMember).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (r *RoomMemberRepository) LeaveRoom(ctx context.Context, roomID, userID string) error {
	if _, err := r.db.NewDelete().Model(&entity.RoomMember{}).Where("room_id = ? AND user_id = ?", roomID, userID).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (r *RoomMemberRepository) DeleteRoomMembers(ctx context.Context, roomID string) error {
	if _, err := r.db.NewDelete().Model(&entity.RoomMember{}).Where("room_id = ?", roomID).Exec(ctx); err != nil {
		return err
	}
	return nil
}

var _ dai.RoomMember = (*RoomMemberRepository)(nil)
