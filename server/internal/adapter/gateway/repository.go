package gateway

import (
	"github.com/murasame29/go-httpserver-template/internal/adapter/gateway/github"
	"github.com/murasame29/go-httpserver-template/internal/adapter/gateway/repository"
	"github.com/uptrace/bun"
)

type Repository struct {
	*repository.UserRepository
	*repository.SessionRepository
	*github.GitHubSerivce
}

func NewRepository(db bun.IDB) *Repository {
	return &Repository{
		UserRepository:    repository.NewUserRepository(db),
		SessionRepository: repository.NewSessionRepository(db),
		GitHubSerivce:     github.NewGitHubSerivce(),
	}
}
