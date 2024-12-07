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
) *WantLearnSkill {
	return &WantLearnSkill{
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

	// TODO: Neptune
	driverRemoteConnection, err := driver.NewNeptuneClient()
		if err != nil {
			return nil, err
		}
		defer driverRemoteConnection.Close()

		g := gremlingo.Traversal_().WithRemote(driverRemoteConnection)

	return s.repo.UpsertWantLearnSkills(ctx, skills)
}
