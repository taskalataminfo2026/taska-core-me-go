package models

import (
	"strings"
	"taska-core-me-go/cmd/api/models"
)

func (*SkillCategoryDb) TableName() string {
	return "skill_categories"
}

type SkillCategoryDb struct {
	ID         int64 `gorm:"column:id;primaryKey"`
	SkillID    int64 `gorm:"column:skill_id"`
	CategoryID int64 `gorm:"column:category_id"`
	IsPrimary  bool  `gorm:"column:is_primary"`
	IsActive   bool  `gorm:"column:is_active"`
}

func (s *SkillCategoryDb) Load(m models.SkillCategory) {
	s.ID = m.ID
	s.SkillID = m.SkillID
	s.CategoryID = m.CategoryID
	s.IsPrimary = m.IsPrimary
	s.IsActive = m.IsActive
}

func (s *SkillCategoryDb) ToDomainModel() models.SkillCategory {
	return models.SkillCategory{
		ID:         s.ID,
		SkillID:    s.SkillID,
		CategoryID: s.CategoryID,
		IsPrimary:  s.IsPrimary,
		IsActive:   s.IsActive,
	}
}

type ParamsSkillsCategorySearchDb struct {
	ID         int64
	SkillID    int64
	CategoryID int64
	IsPrimary  bool
	IsActive   bool
	Limit      int
	Offset     int
}

func (p *ParamsSkillsCategorySearchDb) GetQueryRoles() (string, []interface{}) {
	query := []string{}
	params := []interface{}{}

	if p.ID > 0 {
		query = append(query, "id = ? ")
		params = append(params, p.ID)
	}

	if p.SkillID > 0 {
		query = append(query, "skill_id = ? ")
		params = append(params, p.SkillID)
	}

	if p.CategoryID > 0 {
		query = append(query, "category_id = ? ")
		params = append(params, p.CategoryID)
	}

	if p.IsPrimary == true {
		query = append(query, "is_primary = ? ")
		params = append(params, p.IsPrimary)
	}

	if p.IsActive == true {
		query = append(query, "is_active = ? ")
		params = append(params, p.IsActive)
	}

	return strings.Join(query, " AND "), params
}

func (s *ParamsSkillsCategorySearchDb) ToDomainModel() models.ParamsSkillsCategorySearch {
	return models.ParamsSkillsCategorySearch{
		ID:         s.ID,
		SkillID:    s.SkillID,
		CategoryID: s.CategoryID,
		IsPrimary:  s.IsPrimary,
		IsActive:   s.IsActive,
		Limit:      s.Limit,
		Offset:     s.Offset,
	}
}

func (p *ParamsSkillsCategorySearchDb) ToDB(u *models.ParamsSkillsCategorySearch) {
	p.ID = u.ID
	p.SkillID = u.SkillID
	p.CategoryID = u.CategoryID
	p.IsPrimary = u.IsPrimary
	p.IsActive = u.IsActive
	p.Limit = u.Limit
	p.Offset = u.Offset
}
