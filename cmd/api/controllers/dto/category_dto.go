package dto

import (
	"taska-core-me-go/cmd/api/models"
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
