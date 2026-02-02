package dto

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
	"taska-core-me-go/cmd/api/models"
)

type SkillCategoryDto struct {
	ID         int64 `json:"id"`
	SkillID    int64 `json:"skill_id"`
	CategoryID int64 `json:"category_id"`
	IsPrimary  bool  `json:"is_primary"`
	IsActive   bool  `json:"is_active"`
}

type ParamsSkillsCategorySaveDto struct {
	SkillID    int64 `json:"skill_id"`
	CategoryID int64 `json:"category_id"`
	IsPrimary  bool  `json:"is_primary"`
	IsActive   bool  `json:"is_active"`
}

func (p *ParamsSkillsCategorySaveDto) ToModel() models.ParamsSkillsCategorySave {
	return models.ParamsSkillsCategorySave{
		SkillID:    p.SkillID,
		CategoryID: p.CategoryID,
		IsPrimary:  p.IsPrimary,
		IsActive:   p.IsActive,
	}
}

func SkillCategoryToDto(s models.SkillCategory) SkillCategoryDto {
	return SkillCategoryDto{
		ID:         s.ID,
		SkillID:    s.SkillID,
		CategoryID: s.CategoryID,
		IsPrimary:  s.IsPrimary,
		IsActive:   s.IsActive,
	}
}

type ParamsSkillsCategoryRequestDTO struct {
	ID int64 `json:"id" validate:"required"`
}

func (u *ParamsSkillsCategoryRequestDTO) ParseIDFromParam(c echo.Context) error {
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
