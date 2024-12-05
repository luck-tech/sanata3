package gateway

import (
	"github.com/murasame29/go-httpserver-template/internal/adapter/gateway/github"
	"github.com/murasame29/go-httpserver-template/internal/adapter/gateway/repository"
	"github.com/uptrace/bun"
)

type Repository struct {
	*repository.UserRepository
	*repository.SessionRepository
	*repository.SkillRepository
	*repository.UsedSkillRepository
	*repository.WantLearnSkillRepository
	*github.GitHubSerivce
}

func NewRepository(db bun.IDB) *Repository {
	return &Repository{
		UserRepository:           repository.NewUserRepository(db),
		SessionRepository:        repository.NewSessionRepository(db),
		SkillRepository:          repository.NewSkillRepository(db),
		UsedSkillRepository:      repository.NewUsedSkillRepository(db),
		WantLearnSkillRepository: repository.NewWantLearnSkillRepository(db),
		GitHubSerivce:            github.NewGitHubSerivce(),
	}
}
