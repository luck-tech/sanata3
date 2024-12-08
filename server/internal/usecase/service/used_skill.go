package service

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
)

type UsedSkill struct {
	repo dai.DataAccessInterface
}

func NewUsedSkill(
	repo dai.DataAccessInterface,
) *UsedSkill {
	return &UsedSkill{
		repo: repo,
	}
}

func (s *UsedSkill) UpsertUsedSkill(ctx context.Context, userID string, skill []string) error {
	usedSkill, err := s.repo.GetSkillsByName(ctx, skill)
	if err != nil {
		return err
	}

	var usedSkillIDs []int
	for _, s := range usedSkill {
		usedSkillIDs = append(usedSkillIDs, s.ID)
	}

	var skills []entity.UsedSkill

	for _, s := range usedSkillIDs {
		skills = append(skills, entity.UsedSkill{UserID: userID, SkillID: s})
	}

	return s.repo.UpsertUsedSkills(ctx, userID, skills)
}
