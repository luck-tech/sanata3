package interactor

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/usecase/service"
)

type Skill struct {
	_skill          *service.Skill
	_usedSkill      *service.UsedSkill
	_wantLearnSkill *service.WantLearnSkill
}

func NewSkill(
	skill *service.Skill,
	usedSkill *service.UsedSkill,
	wantLearnSkill *service.WantLearnSkill,
) *Skill {
	return &Skill{
		_skill:          skill,
		_usedSkill:      usedSkill,
		_wantLearnSkill: wantLearnSkill,
	}
}

func (s *Skill) SearchSkill(ctx context.Context, query string, limit int) ([]string, error) {
	skills, err := s._skill.SearchSkills(ctx, query, limit)
	if err != nil {
		return nil, err
	}

	result := make([]string, len(skills))
	for i, skill := range skills {
		result[i] = skill.Name
	}

	return result, nil
}
