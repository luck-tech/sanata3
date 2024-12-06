package dai

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
)

type UsedSkill interface {
	UpsertUsedSkills(ctx context.Context, usedSkill []entity.UsedSkill) error
	GetUsedSkillsByUserID(ctx context.Context, userID string) ([]entity.UsedSkill, error)
}
