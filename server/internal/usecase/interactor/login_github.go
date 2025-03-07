package interactor

import (
	"context"
	"time"

	"github.com/murasame29/go-httpserver-template/internal/framework/jwts"
	"github.com/murasame29/go-httpserver-template/internal/framework/serrors"
	"github.com/murasame29/go-httpserver-template/internal/usecase/service"
)

type Login struct {
	_github    *service.GitHub
	_session   *service.Session
	_user      *service.User
	_skill     *service.Skill
	_usedSkill *service.UsedSkill
	jwt        *jwts.JWTMaker
}

func NewLogin(
	github *service.GitHub,
	session *service.Session,
	user *service.User,
	skill *service.Skill,
	usedSkill *service.UsedSkill,
	jwt *jwts.JWTMaker,
) *Login {
	return &Login{
		_github:    github,
		_session:   session,
		_user:      user,
		_skill:     skill,
		_usedSkill: usedSkill,
		jwt:        jwt,
	}
}

type LoginGitHubParam struct {
	Code string
}

type LoginGithubResult struct {
	JWT       string
	UserID    string
	UserName  string
	Icon      string
	IsNewUser bool
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

	if loginResult.IsNewUser {
		languages, err := i._github.GetUsedLanguage(ctx, loginResult.UserName, token)
		if err != nil {
			return nil, err
		}

		var skills []string
		for k := range languages {
			skills = append(skills, k)
		}

		if err := i._skill.Upsert(ctx, skills); err != nil {
			return nil, err
		}

		if err := i._usedSkill.UpsertUsedSkill(ctx, loginResult.UserID, skills); err != nil {
			return nil, err
		}
	}

	return &LoginGithubResult{
		JWT:       token,
		UserID:    loginResult.UserID,
		UserName:  loginResult.UserName,
		Icon:      loginResult.Icon,
		IsNewUser: loginResult.IsNewUser,
	}, nil
}

type CheckLoginResult struct {
	UserID    string
	SessionID string
}

func (i *Login) CheckLogin(ctx context.Context, token string) (*CheckLoginResult, error) {
	jwtPayload, err := i.jwt.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	session, found, err := i._session.GetSessionByID(ctx, jwtPayload.SessionID)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, serrors.ErrSessionNotFound
	}

	return &CheckLoginResult{
		UserID:    session.UserID,
		SessionID: session.ID,
	}, nil
}
