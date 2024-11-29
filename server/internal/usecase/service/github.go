package service

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
)

type GitHub struct {
	repo dai.DataAccessInterface
}

func NewGitHub(repo dai.DataAccessInterface) *GitHub {
	return &GitHub{
		repo: repo,
	}
}

type LoginGitHubResult struct {
	AccessToken  string
	RefreshToken string
	UserID       string
	UserName     string
	Icon         string
}

func (g *GitHub) Login(ctx context.Context, code string) (*LoginGitHubResult, error) {
	token, err := g.repo.FetchToken(ctx, code)
	if err != nil {
		return nil, err
	}

	userInfo, err := g.repo.GetUserByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	user, found, err := g.repo.GetUser(ctx, string(rune(userInfo.GetID())))
	if err != nil {
		return nil, err
	}
	if !found {
		newUser := &entity.User{
			ID:    string(rune(userInfo.GetID())),
			Email: user.Email,
			Name:  userInfo.GetName(),
			Icon:  userInfo.GetAvatarURL(),
		}
		if err := g.repo.CreateUser(ctx, newUser); err != nil {
			return nil, err
		}
	} else {
		user.Name = userInfo.GetName()
		user.Icon = userInfo.GetAvatarURL()
		if err := g.repo.UpdateUser(ctx, user); err != nil {
			return nil, err
		}
	}

	user, _, err = g.repo.GetUser(ctx, string(rune(userInfo.GetID())))
	if err != nil {
		return nil, err
	}

	return &LoginGitHubResult{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		UserID:       user.ID,
		UserName:     user.Name,
		Icon:         user.Icon,
	}, nil
}
