package models

import (
	"gorm.io/gorm"
	"strings"
	"taska-core-me-go/cmd/api/models"
	"time"
)

func (*SkillsDb) TableName() string {
	return "skills"
}

type ListSkillsDb struct {
	SkillsDb []SkillsDb
}

type SkillsDb struct {
	ID                   int64     `gorm:"column:id"`
	Name                 string    `gorm:"column:name"`
	Slug                 string    `gorm:"column:slug"`
	Description          string    `gorm:"column:description"`
	AvgPriceEstimate     float64   `gorm:"column:avg_price_estimate"`
	RequiresVerification bool      `gorm:"column:requires_verification"`
	RiskLevel            int64     `gorm:"column:risk_level"`
	IsActive             bool      `gorm:"column:is_active"`
	CreatedAt            time.Time `gorm:"column:created_at"`
}

func ToDomainList(list []SkillsDb) []models.Skills {
	result := make([]models.Skills, 0, len(list))
	for i := range list {
		result = append(result, list[i].ToDomainModel())
	}
	return result
}

func (s *SkillsDb) ToDomainModel() models.Skills {
	return models.Skills{
		ID:                   s.ID,
		Name:                 s.Name,
		Slug:                 s.Slug,
		Description:          s.Description,
		AvgPriceEstimate:     s.AvgPriceEstimate,
		RequiresVerification: s.RequiresVerification,
		RiskLevel:            s.RiskLevel,
		IsActive:             s.IsActive,
		CreatedAt:            s.CreatedAt,
	}
}

func (db *SkillsDb) Load(m models.Skills) {
	db.ID = m.ID
	db.Name = m.Name
	db.Slug = m.Slug
	db.Description = m.Description
	db.AvgPriceEstimate = m.AvgPriceEstimate
	db.RequiresVerification = m.RequiresVerification
	db.RiskLevel = m.RiskLevel
	db.IsActive = m.IsActive
	db.CreatedAt = m.CreatedAt
}

type ParamsSkillsSearchDb struct {
	ID                   int64
	Slug                 string
	AvgPriceEstimate     float64
	RequiresVerification bool
	RiskLevel            int64
	IsActive             bool
}

func (p *ParamsSkillsSearchDb) GetQueryRoles() (string, []interface{}) {
	query := []string{}
	params := []interface{}{}

	if p.ID > 0 {
		query = append(query, "id = ? ")
		params = append(params, p.ID)
	}

	if p.Slug > "" {
		query = append(query, "slug = ? ")
		params = append(params, p.Slug)
	}

	if p.AvgPriceEstimate > 0 {
		query = append(query, "avg_price_estimate = ? ")
		params = append(params, p.AvgPriceEstimate)
	}

	if p.RequiresVerification == true {
		query = append(query, "requires_verification = ? ")
		params = append(params, p.RequiresVerification)
	}

	if p.RiskLevel > 0 {
		query = append(query, "risk_level = ? ")
		params = append(params, p.RiskLevel)
	}

	if p.IsActive == true {
		query = append(query, "is_active = ? ")
		params = append(params, p.IsActive)
	}

	return strings.Join(query, " AND "), params
}

func (p *ParamsSkillsSearchDb) ToDB(u *models.ParamsSkillsSearch) {
	p.ID = u.ID
	p.Slug = u.Slug
	p.AvgPriceEstimate = u.AvgPriceEstimate
	p.RequiresVerification = u.RequiresVerification
	p.RiskLevel = u.RiskLevel
	p.IsActive = u.IsActive
}

func (s *SkillsDb) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Local()
	s.CreatedAt = now
	return nil
}
