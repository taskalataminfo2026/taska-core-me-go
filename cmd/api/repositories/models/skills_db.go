package models

import (
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

//func (u *SkillsDb) BeforeCreate(tx *gorm.DB) (err error) {
//	now := time.Now().Local()
//	u.CreatedAt = now
//	return nil
//}
