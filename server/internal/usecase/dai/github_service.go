package dai

import (
	"context"

	"github.com/google/go-github/v67/github"
	"golang.org/x/oauth2"
)

type GitHubService interface {
	FetchToken(ctx context.Context, code string) (*oauth2.Token, error)
	GetUserByToken(ctx context.Context, token *oauth2.Token) (*github.User, error)
}
