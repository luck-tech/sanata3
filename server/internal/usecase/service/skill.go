package service

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
)

type Skill struct {
	repo dai.DataAccessInterface
}

func NewSkill(
	repo dai.DataAccessInterface,
) *Skill {
	return &Skill{
		repo: repo,
	}
}

func (s *Skill) UpsertSkills(ctx context.Context, skills []string) error {
	return s.repo.UpsertSkills(ctx, skills)
}

func (s *Skill) GetSkills(ctx context.Context, skills []int) ([]entity.Skill, error) {
	return s.repo.GetSkills(ctx, skills)
}

func (s *Skill) GetSkillsByName(ctx context.Context, skills []string) ([]entity.Skill, error) {
	return s.repo.GetSkillsByName(ctx, skills)
}

func (s *Skill) SearchSkills(ctx context.Context, query string, limit int) ([]entity.Skill, error) {
	return s.repo.SearchSkills(ctx, query, limit)
}
