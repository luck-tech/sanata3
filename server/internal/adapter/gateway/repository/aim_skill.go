package repository

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
	"github.com/uptrace/bun"
)

type AimSkillRepository struct {
	db bun.IDB
}

func NewAimSkillRepository(db bun.IDB) *AimSkillRepository {
	return &AimSkillRepository{db: db}
}

func (r *AimSkillRepository) UpsertAimSkills(ctx context.Context, aimSkills []entity.AimSkill) error {
	if len(aimSkills) == 0 {
		return nil
	}

	if _, err := r.db.NewDelete().Model(&aimSkills).Where("room_id = ?", aimSkills[0].RoomID).Exec(ctx); err != nil {
		return err
	}

	if _, err := r.db.NewInsert().Model(&aimSkills).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *AimSkillRepository) GetAimSkillsByRoomID(ctx context.Context, roomID string) ([]entity.AimSkill, error) {
	var aimSkills []entity.AimSkill
	if err := r.db.NewSelect().Model(&aimSkills).Where("room_id = ?", roomID).Scan(ctx, &aimSkills); err != nil {
		return nil, err
	}

	return aimSkills, nil
}

func (r *AimSkillRepository) GetAimSkillsByRoomIDs(ctx context.Context, roomIDs []string) ([]entity.AimSkill, error) {
	var aimSkills []entity.AimSkill
	if err := r.db.NewSelect().Model(&aimSkills).Where("room_id IN (?)", bun.In(roomIDs)).Scan(ctx, &aimSkills); err != nil {
		return nil, err
	}

	return aimSkills, nil
}

var _ dai.AimSkill = (*AimSkillRepository)(nil)
