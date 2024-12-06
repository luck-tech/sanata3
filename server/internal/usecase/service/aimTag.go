package service

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
)

type AimSkill struct {
	repo dai.DataAccessInterface
}

func NewAimSkill(
	repo dai.DataAccessInterface,
) *AimSkill {
	return &AimSkill{
		repo: repo,
	}
}

func (s *AimSkill) List(ctx context.Context, roomID string) ([]entity.AimSkill, error) {
	return s.repo.GetAimSkillsByRoomID(ctx, roomID)
}

func (s *AimSkill) ListByRoomIDs(ctx context.Context, roomIDs []string) (map[string][]entity.AimSkill, []int, error) {
	aimSkills, err := s.repo.GetAimSkillsByRoomIDs(ctx, roomIDs)
	if err != nil {
		return nil, nil, err
	}

	result := make(map[string][]entity.AimSkill)
	skillIDs := make(map[int]struct{})
	for _, aimSkill := range aimSkills {
		result[aimSkill.RoomID] = append(result[aimSkill.RoomID], aimSkill)
		skillIDs[aimSkill.SkillID] = struct{}{}
	}

	var skills []int
	for skillID := range skillIDs {
		skills = append(skills, skillID)
	}

	return result, skills, nil
}

func (s *AimSkill) Upsert(ctx context.Context, roomID string, skills []string) error {
	aimSkills, err := s.repo.GetSkillsByName(ctx, skills)
	if err != nil {
		return err
	}

	var aimSkillIDs []int
	for _, s := range aimSkills {
		aimSkillIDs = append(aimSkillIDs, s.ID)
	}

	var newAimSkills []entity.AimSkill

	for _, s := range aimSkillIDs {
		newAimSkills = append(newAimSkills, entity.AimSkill{RoomID: roomID, SkillID: s})
	}

	return s.repo.UpsertAimSkills(ctx, newAimSkills)
}
