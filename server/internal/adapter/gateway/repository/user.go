package repository

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db bun.IDB
}

func NewUserRepository(db bun.IDB) *UserRepository {
	return &UserRepository{db: db}
}

func (s *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	panic("impl me")
}

func (s *UserRepository) GetUser(ctx context.Context, id string) (*entity.User, bool, error) {
	panic("impl me")
}

func (s *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	panic("impl me")
}

func (s *UserRepository) DeleteUser(ctx context.Context, id string) error {
	panic("impl me")
}

var _ dai.User = (*UserRepository)(nil)
