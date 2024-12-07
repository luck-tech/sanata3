package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
	"github.com/uptrace/bun"
)

type RoomRepository struct {
	db bun.IDB
}

func NewRoomRepository(db bun.IDB) *RoomRepository {
	return &RoomRepository{db: db}
}

func (s *RoomRepository) CreateRoom(ctx context.Context, room *entity.Room) error {
	_, err := s.db.NewInsert().
		Model(room).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoomRepository) UpdateRoom(ctx context.Context, room *entity.Room) error {
	_, err := s.db.NewUpdate().
		Model(room).
		WherePK().
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoomRepository) GetRoom(ctx context.Context, roomID string) (*entity.Room, bool, error) {
	var room entity.Room
	err := s.db.NewSelect().
		Model(&room).
		Where("id = ?", roomID).
		Scan(ctx, &room)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, false, err
		}
		return nil, false, nil
	}

	return &room, true, nil
}

func (s *RoomRepository) ListRoom(ctx context.Context, userID string) ([]entity.Room, error) {
	var rooms []entity.Room
	// SELECT * FROM rooms JOIN
	err := s.db.NewSelect().Model(&rooms).Join(
		"LEFT JOIN room_members ON room.id = room_members.room_id",
	).Where("room_members.user_id = ?", userID).Scan(ctx, &rooms)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (s *RoomRepository) DeleteRoom(ctx context.Context, userID string) error {
	if _, err := s.db.NewDelete().Model(&entity.Room{}).Where("id = ?", userID).Exec(ctx); err != nil {
		return err
	}

	return nil
}

var _ dai.Room = (*RoomRepository)(nil)
