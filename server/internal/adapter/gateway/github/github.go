package github

import (
	"context"

	v67 "github.com/google/go-github/v67/github"
	"github.com/murasame29/go-httpserver-template/cmd/config"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type GitHubSerivce struct {
	oac oauth2.Config
}

func NewGitHubSerivce() *GitHubSerivce {
	return &GitHubSerivce{
		oac: oauth2.Config{
			ClientID:     config.Config.GitHub.ClientID,
			ClientSecret: config.Config.GitHub.ClientSecret,
			Scopes:       []string{},
			Endpoint: oauth2.Endpoint{
				AuthURL:  github.Endpoint.AuthURL,
				TokenURL: github.Endpoint.TokenURL,
			},
			RedirectURL: config.Config.GitHub.RedirectURI,
		},
	}
}

func (s *GitHubSerivce) FetchToken(ctx context.Context, code string) (*oauth2.Token, error) {
	panic("impl me")
}

func (s *GitHubSerivce) GetUserByToken(ctx context.Context, token *oauth2.Token) (*v67.User, error) {
	panic("impl me")
}

var _ dai.GitHubService = (*GitHubSerivce)(nil)
