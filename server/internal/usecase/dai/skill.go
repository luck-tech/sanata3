package dai

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
)

type Skill interface {
	UpsertSkills(ctx context.Context, skills []string) error
	GetSkills(ctx context.Context, skills []int) ([]entity.Skill, error)
	GetSkillsByName(ctx context.Context, skills []string) ([]entity.Skill, error)
	SearchSkills(ctx context.Context, query string, limit int) ([]entity.Skill, error)
}
