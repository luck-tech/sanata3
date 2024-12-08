package repository

import (
	"context"
	"log"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
	"github.com/uptrace/bun"
)

type WantLearnSkillRepository struct {
	db bun.IDB
}

func NewWantLearnSkillRepository(db bun.IDB) *WantLearnSkillRepository {
	return &WantLearnSkillRepository{db: db}
}

func (r *WantLearnSkillRepository) UpsertWantLearnSkills(ctx context.Context, userID string, wantLearnSkills []entity.WantLearnSkill) error {
	if _, err := r.db.NewDelete().Model(&wantLearnSkills).Where("user_id = ?", userID).Exec(ctx); err != nil {
		return err
	}

	skills, err := r.GetWantLearnSkills(ctx, userID)
	if err != nil {
		return err
	}
	log.Println("fugafuga1", skills, userID)

	if len(wantLearnSkills) == 0 {
		return nil
	}

	if _, err := r.db.NewInsert().Model(&wantLearnSkills).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *WantLearnSkillRepository) GetWantLearnSkills(ctx context.Context, userID string) ([]entity.WantLearnSkill, error) {
	wantLearnSkill := []entity.WantLearnSkill{}
	if err := r.db.NewSelect().Model(&wantLearnSkill).Where("user_id = ?", userID).Scan(ctx, &wantLearnSkill); err != nil {
		return nil, err
	}

	return wantLearnSkill, nil
}

var _ dai.WantLearnSkill = (*WantLearnSkillRepository)(nil)
