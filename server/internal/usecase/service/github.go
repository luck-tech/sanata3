package service

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/framework/contexts"
	"github.com/murasame29/go-httpserver-template/internal/framework/serrors"
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

	userInfo, err := g.repo.GetUserByToken(ctx, token.AccessToken)
	if err != nil {
		return nil, err
	}

	user, found, err := g.repo.GetUser(ctx, string(rune(userInfo.ID)))
	if err != nil {
		return nil, err
	}
	if !found {
		newUser := &entity.User{
			ID:    string(rune(userInfo.ID)),
			Email: user.Email,
			Name:  userInfo.Login,
			Icon:  userInfo.AvatarURL,
		}
		if err := g.repo.CreateUser(ctx, newUser); err != nil {
			return nil, err
		}
	} else {
		user.Name = userInfo.Login
		user.Icon = userInfo.AvatarURL
		if err := g.repo.UpdateUser(ctx, user); err != nil {
			return nil, err
		}
	}

	user, _, err = g.repo.GetUser(ctx, string(rune(userInfo.ID)))
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

func (g *GitHub) GetUsedLanguage(ctx context.Context) (map[string]int, error) {
	sessionID := contexts.GetSessionID(ctx)
	session, found, err := g.repo.GetSessionByID(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, serrors.ErrSessionNotFound
	}

	userID := contexts.GetUserID(ctx)
	user, found, err := g.repo.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, serrors.ErrSessionNotFound
	}

	return g.repo.GetUserUseLanguagesByID(ctx, session.AccessToken, user.Name)
}
