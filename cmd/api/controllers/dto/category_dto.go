package dto

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/utils"
	"time"
)

type CategoryDto struct {
	ID          int64     `json:"id"`
	RootID      int64     `json:"root_id,omitempty"`
	ParentID    int64     `json:"parent_id,omitempty"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description,omitempty"`
	Icon        string    `json:"icon,omitempty"`
	IsActive    bool      `json:"is_active"`
	SortOrder   int       `json:"sort_order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (c *CategoryDto) FromModel(modelsList []models.Category) []CategoryDto {
	items := make([]CategoryDto, 0, len(modelsList))

	for _, skill := range modelsList {
		items = append(items, CategoryToDto(skill))
	}

	return items
}

func CategoryToDto(s models.Category) CategoryDto {
	return CategoryDto{
		ID:          s.ID,
		RootID:      s.RootID,
		ParentID:    s.ParentID,
		Name:        s.Name,
		Slug:        s.Slug,
		Description: s.Description,
		Icon:        s.Icon,
		IsActive:    s.IsActive,
		SortOrder:   s.SortOrder,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

type ParamsCategorySearchDto struct {
	ID        int64  `json:"id"`
	RootID    int64  `json:"root_id"`
	ParentID  int64  `json:"parent_id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Icon      string `json:"icon"`
	IsActive  bool   `json:"is_active"`
	SortOrder int    `json:"sort_order"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
}

func (p *ParamsCategorySearchDto) BindCategorySearchFilter(c echo.Context) error {

	if err := utils.GetInt64Query(c, "id", &p.ID); err != nil {
		return err
	}

	if err := utils.GetInt64Query(c, "root_id", &p.ID); err != nil {
		return err
	}

	if err := utils.GetInt64Query(c, "parent_id", &p.ID); err != nil {
		return err
	}

	if v := c.QueryParam("name"); v != "" {
		p.Slug = v
	}

	if v := c.QueryParam("slug"); v != "" {
		p.Slug = v
	}

	if v := c.QueryParam("icon"); v != "" {
		p.Slug = v
	}

	if err := utils.GetBoolQuery(c, "is_active", &p.IsActive); err != nil {
		return err
	}

	if err := utils.GetIntQuery(c, "sort_order", &p.Limit); err != nil {
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

func (p *ParamsCategorySearchDto) ToModel() models.ParamsCategorySearch {
	return models.ParamsCategorySearch{
		ID:        p.ID,
		RootID:    p.RootID,
		ParentID:  p.ParentID,
		Name:      utils.StrinToLower(p.Name),
		Slug:      utils.StrinToLower(p.Slug),
		Icon:      p.Icon,
		IsActive:  p.IsActive,
		SortOrder: p.SortOrder,
		Limit:     p.Limit,
		Offset:    p.Offset,
	}
}

type ParamsCategorySaveDto struct {
	RootID      int64  `json:"root_id"`
	ParentID    int64  `json:"parent_id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	IsActive    bool   `json:"is_active"`
	SortOrder   int    `json:"sort_order"`
}

func (p *ParamsCategorySaveDto) ToModel() models.ParamsCategorySave {
	return models.ParamsCategorySave{
		RootID:      p.RootID,
		ParentID:    p.ParentID,
		Name:        utils.StrinToLower(p.Name),
		Slug:        utils.StrinToLower(p.Slug),
		Description: p.Description,
		Icon:        p.Icon,
		IsActive:    p.IsActive,
		SortOrder:   p.SortOrder,
	}
}

type ParamsCategoryRequestDTO struct {
	ID int64 `json:"id" validate:"required"`
}

func (u *ParamsCategoryRequestDTO) ParseIDFromParam(c echo.Context) error {
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
