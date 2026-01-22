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

type ParamsSkillsSearchDto struct {
	ID                   int64   `json:"id"`
	Slug                 string  `json:"slug"`
	AvgPriceEstimate     float64 `json:"avg_price_estimate"`
	RequiresVerification bool    `json:"requires_verification"`
	RiskLevel            int64   `json:"risk_level"`
	IsActive             bool    `json:"is_active"`
	Limit                int     `json:"limit"`
	Offset               int     `json:"offset"`
}

func (u *ParamsSkillsSearchDto) BindSkillsSearchFilter(c echo.Context) error {

	if err := utils.GetInt64Query(c, "id", &u.ID); err != nil {
		return err
	}

	if v := c.QueryParam("slug"); v != "" {
		u.Slug = v
	}

	if err := utils.GetFloat64Query(c, "avg_price_estimate", &u.AvgPriceEstimate); err != nil {
		return err
	}

	if err := utils.GetBoolQuery(c, "requires_verification", &u.RequiresVerification); err != nil {
		return err
	}

	if err := utils.GetInt64Query(c, "risk_level", &u.RiskLevel); err != nil {
		return err
	}

	if err := utils.GetBoolQuery(c, "is_active", &u.IsActive); err != nil {
		return err
	}

	if err := utils.GetIntQuery(c, "limit", &u.Limit); err != nil {
		return err
	}

	if err := utils.GetIntQuery(c, "offset", &u.Offset); err != nil {
		return err
	}

	if u.Limit == 0 {
		u.Limit = 20
	}

	return nil
}

func (v *ParamsSkillsSearchDto) ToModel() models.ParamsSkillsSearch {
	return models.ParamsSkillsSearch{
		ID:                   v.ID,
		Slug:                 utils.StrinToLower(v.Slug),
		AvgPriceEstimate:     v.AvgPriceEstimate,
		RequiresVerification: v.RequiresVerification,
		RiskLevel:            v.RiskLevel,
		IsActive:             v.IsActive,
		Limit:                v.Limit,
	}
}
