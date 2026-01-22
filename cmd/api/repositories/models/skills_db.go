package models

import (
	"strings"
	"taska-core-me-go/cmd/api/models"
	"time"
)

func (SkillsDb) TableName() string { return "skills" }

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

func ToDomainList(list []SkillsDb) []models.SkillsResponse {
	skills := make([]models.SkillsResponse, 0, len(list))

	for _, skill := range list {
		skills = append(skills, skill.ToDomainModel())
	}
	return skills
}

func (s SkillsDb) ToDomainModel() models.SkillsResponse {
	return models.SkillsResponse{
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

type ParamsSkillsSearchDb struct {
	ID                   int64
	Slug                 string
	AvgPriceEstimate     float64
	RequiresVerification bool
	RiskLevel            int64
	IsActive             bool
}

func (u *ParamsSkillsSearchDb) GetQueryRoles() (string, []interface{}) {
	query := []string{}
	params := []interface{}{}

	if u.ID > 0 {
		query = append(query, "id = ? ")
		params = append(params, u.ID)
	}

	if u.Slug > "" {
		query = append(query, "slug = ? ")
		params = append(params, u.Slug)
	}

	if u.AvgPriceEstimate > 0 {
		query = append(query, "avg_price_estimate = ? ")
		params = append(params, u.AvgPriceEstimate)
	}

	if u.RequiresVerification == true {
		query = append(query, "requires_verification = ? ")
		params = append(params, u.RequiresVerification)
	}

	if u.RiskLevel > 0 {
		query = append(query, "risk_level = ? ")
		params = append(params, u.RiskLevel)
	}

	if u.IsActive == true {
		query = append(query, "is_active = ? ")
		params = append(params, u.IsActive)
	}

	return strings.Join(query, " AND "), params
}

func (db *ParamsSkillsSearchDb) ToDB(u *models.ParamsSkillsSearch) {
	db.ID = u.ID
	db.Slug = u.Slug
	db.AvgPriceEstimate = u.AvgPriceEstimate
	db.RequiresVerification = u.RequiresVerification
	db.RiskLevel = u.RiskLevel
	db.IsActive = u.IsActive
}
