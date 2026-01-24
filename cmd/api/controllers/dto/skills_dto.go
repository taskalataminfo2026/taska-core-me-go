package dto

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/utils"
	"time"
)

type SkillsDto struct {
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

func (s *SkillsDto) FromModel(modelsList []models.Skills) []SkillsDto {
	items := make([]SkillsDto, 0, len(modelsList))

	for _, skill := range modelsList {
		items = append(items, SkillToDto(skill))
	}

	return items
}

func SkillToDto(s models.Skills) SkillsDto {
	return SkillsDto{
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

func (p *ParamsSkillsSearchDto) BindSkillsSearchFilter(c echo.Context) error {

	if err := utils.GetInt64Query(c, "id", &p.ID); err != nil {
		return err
	}

	if v := c.QueryParam("slug"); v != "" {
		p.Slug = v
	}

	if err := utils.GetFloat64Query(c, "avg_price_estimate", &p.AvgPriceEstimate); err != nil {
		return err
	}

	if err := utils.GetBoolQuery(c, "requires_verification", &p.RequiresVerification); err != nil {
		return err
	}

	if err := utils.GetInt64Query(c, "risk_level", &p.RiskLevel); err != nil {
		return err
	}

	if err := utils.GetBoolQuery(c, "is_active", &p.IsActive); err != nil {
		return err
	}

	if err := utils.GetIntQuery(c, "limit", &p.Limit); err != nil {
		return err
	}

	if err := utils.GetIntQuery(c, "offset", &p.Offset); err != nil {
		return err
	}

	if p.Limit == 0 {
		p.Limit = 20
	}

	return nil
}

func (p *ParamsSkillsSearchDto) ToModel() models.ParamsSkillsSearch {
	return models.ParamsSkillsSearch{
		ID:                   p.ID,
		Slug:                 utils.StrinToLower(p.Slug),
		AvgPriceEstimate:     p.AvgPriceEstimate,
		RequiresVerification: p.RequiresVerification,
		RiskLevel:            p.RiskLevel,
		IsActive:             p.IsActive,
		Limit:                p.Limit,
	}
}

type ParamsSkillsUpsertDto struct {
	Name                 string  `json:"name" validate:"required"`
	Slug                 string  `json:"slug"  validate:"required"`
	Description          string  `json:"description" validate:"required"`
	AvgPriceEstimate     float64 `json:"avg_price_estimate" validate:"required"`
	RequiresVerification bool    `json:"requires_verification" validate:"required"`
	RiskLevel            int64   `json:"risk_level" validate:"required"`
	IsActive             bool    `json:"is_active" validate:"required"`
}

func (p *ParamsSkillsUpsertDto) Bind(c echo.Context) error {
	if err := c.Bind(p); err != nil {
		return err
	}
	return validator.New().Struct(p)
}

func (p *ParamsSkillsUpsertDto) ToModel() models.ParamsSkillsSave {
	return models.ParamsSkillsSave{
		Name:                 utils.StrinToLower(p.Name),
		Slug:                 utils.StrinToLower(p.Slug),
		Description:          p.Description,
		AvgPriceEstimate:     p.AvgPriceEstimate,
		RequiresVerification: p.RequiresVerification,
		RiskLevel:            p.RiskLevel,
		IsActive:             p.IsActive,
	}
}

type ParamsSkillsRequestDTO struct {
	ID int64 `json:"id" validate:"required"`
}

func (u *ParamsSkillsRequestDTO) ParseIDFromParam(c echo.Context) error {
	idStr := strings.TrimSpace(c.Param("id"))
	if idStr == "" {
		return fmt.Errorf("el parámetro 'id' es obligatorio")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fmt.Errorf("id inválido: debe ser un número entero")
	}

	if id <= 0 {
		return fmt.Errorf("id inválido: debe ser mayor que cero")
	}

	u.ID = id
	return nil
}
