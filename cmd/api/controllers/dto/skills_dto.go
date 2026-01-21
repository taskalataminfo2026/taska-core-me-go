package dto

import (
	"github.com/labstack/echo/v4"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/utils"
	"time"
)

type SkillsResponseDto struct {
	ID                   int64     `json:"id"`
	Name                 string    `json:"name"`
	Slug                 string    `json:"slug"`
	Description          string    `json:"description"`
	AvgPriceEstimate     float64   `json:"avg_price_estimate"`
	RequiresVerification bool      `json:"requires_verification"`
	RiskLevel            int64     `json:"risk_level"`
	IsActive             bool      `json:"is_active"`
	CreatedAt            time.Time `json:"created_at"`
}

func (s *SkillsResponseDto) FromModel(modelsList []models.SkillsResponse) []SkillsResponseDto {
	items := make([]SkillsResponseDto, 0, len(modelsList))

	for _, skill := range modelsList {
		items = append(items, SkillToDto(skill))
	}

	return items
}

func SkillToDto(s models.SkillsResponse) SkillsResponseDto {
	return SkillsResponseDto{
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

type SkillsRequestDto struct {
	ID                   int64   `json:"id"`
	Name                 string  `json:"name"`
	Slug                 string  `json:"slug"`
	AvgPriceEstimate     float64 `json:"avg_price_estimate"`
	RequiresVerification bool    `json:"requires_verification"`
	RiskLevel            int64   `json:"risk_level"`
	IsActive             bool    `json:"is_active"`
}

func (u *SkillsRequestDto) Bind(c echo.Context) error {
	if err := c.Bind(u); err != nil {
		return err
	}
	return nil
}

func (v *SkillsRequestDto) ToModel() models.SkillsRequest {
	return models.SkillsRequest{
		ID:                   v.ID,
		Name:                 utils.StrinToLower(v.Name),
		Slug:                 utils.StrinToLower(v.Slug),
		AvgPriceEstimate:     v.AvgPriceEstimate,
		RequiresVerification: v.RequiresVerification,
		RiskLevel:            v.RiskLevel,
		IsActive:             v.IsActive,
	}
}
