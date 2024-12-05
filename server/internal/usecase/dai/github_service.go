package dai

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"golang.org/x/oauth2"
)

type GitHubService interface {
	FetchToken(ctx context.Context, code string) (*oauth2.Token, error)
	GetUserByToken(ctx context.Context, accessToken string) (*entity.GitHubUser, error)
	GetUserUseLanguagesByID(ctx context.Context, accessToken, username string) (map[string]int, error)
}
