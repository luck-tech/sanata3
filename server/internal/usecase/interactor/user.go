package interactor

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/service"
)

type User struct {
	_session        *service.Session
	_user           *service.User
	_skill          *service.Skill
	_usedSkill      *service.UsedSkill
	_wantLearnSkill *service.WantLearnSkill
}

func NewUser(
	session *service.Session,
	user *service.User,
	skill *service.Skill,
	usedSkill *service.UsedSkill,
	wantLearnSkill *service.WantLearnSkill,
) *User {
	return &User{
		_session:        session,
		_user:           user,
		_skill:          skill,
		_usedSkill:      usedSkill,
		_wantLearnSkill: wantLearnSkill,
	}
}

func (u *User) GetUser(ctx context.Context, id string) (*entity.User, []entity.Skill, []entity.Skill, error) {
	user, found, err := u._user.GetUser(ctx, id)
	if err != nil {
		return nil, nil, nil, err
	}

	if !found {
		return nil, nil, nil, err
	}

	var usedSkillIDs []int
	var wantLeanSkillIDs []int

	for _, us := range user.UsedSkill {
		usedSkillIDs = append(usedSkillIDs, us.SkillID)
	}

	for _, wls := range user.WantLeanSkills {
		wantLeanSkillIDs = append(wantLeanSkillIDs, wls.SkillID)
	}

	usedSkills, err := u._skill.GetSkills(ctx, usedSkillIDs)
	if err != nil {
		return nil, nil, nil, err
	}

	wantLeanSkills, err := u._skill.GetSkills(ctx, wantLeanSkillIDs)
	if err != nil {
		return nil, nil, nil, err
	}

	return user.User, usedSkills, wantLeanSkills, nil
}

type UpdateUserParam struct {
	UserID          string
	Description     string
	UsedSkills      []string
	WantLearnSkills []string
}

func (u *User) UpdateUser(ctx context.Context, param UpdateUserParam) (*entity.User, []entity.Skill, []entity.Skill, error) {
	if err := u._user.UpdateUser(ctx, &entity.User{
		ID:          param.UserID,
		Description: param.Description,
	}); err != nil {
		return nil, nil, nil, err
	}

	if err := u._skill.UpsertSkills(ctx, param.UsedSkills); err != nil {
		return nil, nil, nil, err
	}

	if err := u._skill.UpsertSkills(ctx, param.WantLearnSkills); err != nil {
		return nil, nil, nil, err
	}

	if err := u._usedSkill.UpsertUsedSkill(ctx, param.UserID, param.UsedSkills); err != nil {
		return nil, nil, nil, err
	}

	if err := u._wantLearnSkill.UpsertWantLearnSkill(ctx, param.UserID, param.WantLearnSkills); err != nil {
		return nil, nil, nil, err
	}

	return u.GetUser(ctx, param.UserID)
}
