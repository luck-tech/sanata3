package service

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
)

type User struct {
	repo dai.DataAccessInterface
}

func NewUser(repo dai.DataAccessInterface) *User {
	return &User{
		repo: repo,
	}
}

func (u *User) Update(ctx context.Context, user *entity.User) error {
	return u.repo.UpdateUser(ctx, user)
}

func (u *User) Get(ctx context.Context, userID string) (*entity.User, bool, error) {
	return u.repo.GetUser(ctx, userID)
}

func (u *User) GetInfo(ctx context.Context, id string) (*entity.UserInfo, bool, error) {
	user, found, err := u.repo.GetUser(ctx, id)
	if err != nil {
		return nil, false, err
	}

	if !found {
		return nil, false, nil
	}

	usedSkill, err := u.repo.GetUsedSkillsByUserID(ctx, id)
	if err != nil {
		return nil, false, err
	}

	wantLearnSkill, err := u.repo.GetWantLearnSkills(ctx, id)
	if err != nil {
		return nil, false, err
	}

	return &entity.UserInfo{
		User:           user,
		UsedSkill:      usedSkill,
		WantLeanSkills: wantLearnSkill,
	}, true, nil
}

func (s *User) List(ctx context.Context, userIDs []string) ([]entity.User, error) {
	return s.repo.GetUsers(ctx, userIDs)
}
