package service

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
)

type WantLearnSkill struct {
	repo dai.DataAccessInterface
}

func NewWantLearnSkill(
	repo dai.DataAccessInterface,
) *UsedSkill {
	return &UsedSkill{
		repo: repo,
	}
}

func (s *WantLearnSkill) UpsertWantLearnSkill(ctx context.Context, userID string, skill []string) error {
	wantLearnSkill, err := s.repo.GetSkillsByName(ctx, skill)
	if err != nil {
		return err
	}

	var WantLearnSkillIDs []int
	for _, s := range wantLearnSkill {
		WantLearnSkillIDs = append(WantLearnSkillIDs, s.ID)
	}

	var skills []entity.WantLearnSkill
	for _, s := range WantLearnSkillIDs {
		skills = append(skills, entity.WantLearnSkill{UserID: userID, SkillID: s})
	}

	return s.repo.UpsertWantLearnSkills(ctx, skills)
}
