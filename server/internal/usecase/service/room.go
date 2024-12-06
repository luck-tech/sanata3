package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
)

type Room struct {
	repo dai.DataAccessInterface
}

func NewRoom(
	repo dai.DataAccessInterface,
) *Room {
	return &Room{
		repo: repo,
	}
}

func (s *Room) Get(ctx context.Context, roomID string) (*entity.Room, error) {
	room, found, err := s.repo.GetRoom(ctx, roomID)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, nil
	}

	return room, nil
}

func (s *Room) List(ctx context.Context, userID string) ([]entity.Room, error) {
	return s.repo.ListRoom(ctx, userID)
}

func (s *Room) Create(ctx context.Context, name, description, ownerID string) (string, error) {
	roomID := uuid.NewString()
	if err := s.repo.CreateRoom(ctx, &entity.Room{
		ID:          roomID,
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
	}); err != nil {
		return "", err
	}

	return roomID, nil
}

func (s *Room) Update(ctx context.Context, room *entity.Room) error {
	return s.repo.UpdateRoom(ctx, room)
}

func (s *Room) Delete(ctx context.Context, roomID string) error {
	return s.repo.DeleteRoom(ctx, roomID)
}
