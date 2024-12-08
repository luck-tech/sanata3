package repository

import (
	"context"
	"database/sql"
	"errors"

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
	_, err := s.db.NewInsert().
		Model(user).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserRepository) GetUser(ctx context.Context, id string) (*entity.User, bool, error) {
	var user entity.User
	err := s.db.NewSelect().
		Model(&user).
		Where("id = ?", id).
		Scan(ctx, &user)

	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, false, err
		}
		return nil, false, nil
	}

	return &user, true, nil
}

func (s *UserRepository) GetUsers(ctx context.Context, userIDs []string) ([]entity.User, error) {
	users := []entity.User{}
	err := s.db.NewSelect().
		Model(&users).
		Where("id IN (?)", bun.In(userIDs)).
		Scan(ctx, &users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	if _, err := s.db.NewUpdate().Model(user).WherePK().Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (s *UserRepository) DeleteUser(ctx context.Context, id string) error {
	if _, err := s.db.NewDelete().Model(&entity.User{}).Where("id = ?", id).Exec(ctx); err != nil {
		return err
	}

	return nil
}

var _ dai.User = (*UserRepository)(nil)
