package dai

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
)

type AimSkill interface {
	UpsertAimSkills(ctx context.Context, aimSkills []entity.AimSkill) error
	GetAimSkillsByRoomID(ctx context.Context, roomID string) ([]entity.AimSkill, error)
	GetAimSkillsByRoomIDs(ctx context.Context, roomIDs []string) ([]entity.AimSkill, error)
}
