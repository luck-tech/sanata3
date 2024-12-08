package repository

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
	"github.com/uptrace/bun"
)

type UsedSkillRepository struct {
	db bun.IDB
}

func NewUsedSkillRepository(db bun.IDB) *UsedSkillRepository {
	return &UsedSkillRepository{db: db}
}

func (r *UsedSkillRepository) UpsertUsedSkills(ctx context.Context, userID string, usedSkill []entity.UsedSkill) error {
	if _, err := r.db.NewDelete().Model(&usedSkill).Where("user_id = ?", userID).Exec(ctx); err != nil {
		return err
	}

	if len(usedSkill) == 0 {
		return nil
	}

	if _, err := r.db.NewInsert().Model(&usedSkill).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *UsedSkillRepository) GetUsedSkillsByUserID(ctx context.Context, userID string) ([]entity.UsedSkill, error) {
	usedSkill := []entity.UsedSkill{}
	if err := r.db.NewSelect().Model(&usedSkill).Where("user_id = ?", userID).Scan(ctx, &usedSkill); err != nil {
		return nil, err
	}

	return usedSkill, nil
}

var _ dai.UsedSkill = (*UsedSkillRepository)(nil)
