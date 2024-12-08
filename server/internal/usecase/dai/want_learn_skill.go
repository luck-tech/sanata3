package dai

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
)

type WantLearnSkill interface {
	UpsertWantLearnSkills(ctx context.Context, userID string, wantLearnSkills []entity.WantLearnSkill) error
	GetWantLearnSkills(ctx context.Context, userID string) ([]entity.WantLearnSkill, error)
}
