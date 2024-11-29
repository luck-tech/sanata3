package interactor

import (
	"context"
	"time"

	"github.com/murasame29/go-httpserver-template/internal/framework/jwts"
	"github.com/murasame29/go-httpserver-template/internal/usecase/service"
)

type Login struct {
	_github  *service.GitHub
	_session *service.Session
	jwt      *jwts.JWTMaker
}

func NewLogin(
	github *service.GitHub,
	session *service.Session,
	jwt *jwts.JWTMaker,
) *Login {
	return &Login{
		_github:  github,
		_session: session,
		jwt:      jwt,
	}
}

type LoginGitHubParam struct {
	Code string
}

type LoginGithubResult struct {
	JWT      string
	UserID   string
	UserName string
	Icon     string
}

func (i *Login) GitHub(ctx context.Context, param LoginGitHubParam) (*LoginGithubResult, error) {
	loginResult, err := i._github.Login(ctx, param.Code)
	if err != nil {
		return nil, err
	}

	sessionID, err := i._session.UpsertSession(ctx, loginResult.UserID, loginResult.AccessToken, loginResult.RefreshToken)
	if err != nil {
		return nil, err
	}

	token, err := i.jwt.CreateToken(sessionID, time.Hour*24*30)
	if err != nil {
		return nil, err
	}

	return &LoginGithubResult{
		JWT:      token,
		UserID:   loginResult.UserID,
		UserName: loginResult.UserName,
		Icon:     loginResult.Icon,
	}, nil
}
