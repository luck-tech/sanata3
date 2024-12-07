package repository

import (
	"context"
	"fmt"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/usecase/dai"
	"github.com/uptrace/bun"
)

type SkillRepository struct {
	db bun.IDB
}

func NewSkillRepository(db bun.IDB) *SkillRepository {
	return &SkillRepository{db: db}
}

func (r *SkillRepository) UpsertSkills(ctx context.Context, skills []string) error {
	if len(skills) == 0 {
		return nil // Nothing to do
	}

	newSkills := make([]entity.Skill, 0, len(skills))
	for _, skillName := range skills {
		newSkills = append(newSkills, entity.Skill{Name: skillName})
	}

	_, err := r.db.NewInsert().
		Model(&newSkills).
		On("conflict (name) do nothing").
		Exec(ctx)

	if err != nil {
		return fmt.Errorf("upserting skills: %w", err)
	}

	return nil
}

func (r *SkillRepository) GetSkills(ctx context.Context, skills []int) ([]entity.Skill, error) {
	if len(skills) == 0 {
		return nil, nil // Nothing to do
	}

	var skillList []entity.Skill
	err := r.db.NewSelect().Model(&skillList).Where("id IN (?)", bun.In(skills)).Scan(ctx, &skillList)
	if err != nil {
		return nil, err
	}

	return skillList, nil
}

func (r *SkillRepository) GetSkillsByName(ctx context.Context, skills []string) ([]entity.Skill, error) {
	if len(skills) == 0 {
		return nil, nil // Nothing to do
	}

	var skillList []entity.Skill
	err := r.db.NewSelect().Model(&skillList).Where("name IN (?)", bun.In(skills)).Scan(ctx, &skillList)
	if err != nil {
		return nil, err
	}

	return skillList, nil
}

func (r *SkillRepository) SearchSkills(ctx context.Context, query string, limit int) ([]entity.Skill, error) {
	var skillList []entity.Skill

	// Subquery to count used skills.
	usedSkillsCount := r.db.NewSelect().
		TableExpr("used_skills AS uc").
		ColumnExpr("skill_id, COUNT(*) AS used_count").
		Group("skill_id")

	// Subquery to count wanted skills (if needed).  Comment out if not required.
	wantLearnSkillsCount := r.db.NewSelect().
		TableExpr("want_learn_skills AS wlc").
		ColumnExpr("skill_id, COUNT(*) AS want_learn_count").
		Group("skill_id")

	q := r.db.NewSelect().
		Model(&skillList).
		Join("JOIN (?) AS uc ON skill.id = uc.skill_id", usedSkillsCount).
		Join("LEFT JOIN (?) AS wlc ON skill.id = wlc.skill_id", wantLearnSkillsCount)

	if len(query) != 0 {
		q = q.Where("skill.name LIKE ?", "%"+query+"%")
	}

	if err := q.Order("uc.used_count DESC").Order("wlc.want_learn_count DESC").Limit(limit).Scan(ctx, &skillList); err != nil {
		return nil, err
	}

	return skillList, nil
}

var _ dai.Skill = (*SkillRepository)(nil)
