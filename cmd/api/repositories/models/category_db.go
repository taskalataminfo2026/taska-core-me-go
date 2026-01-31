package models

import (
	"gorm.io/gorm"
	"strings"
	"taska-core-me-go/cmd/api/models"
	"time"
)

func (*CategoryDb) TableName() string {
	return "categories"
}

type CategoryDb struct {
	ID          int64     `gorm:"column:id;primaryKey"`
	RootID      int64     `gorm:"column:root_id"`
	ParentID    int64     `gorm:"column:parent_id"`
	Name        string    `gorm:"column:name"`
	Slug        string    `gorm:"column:slug"`
	Description string    `gorm:"column:description"`
	Icon        string    `gorm:"column:icon"`
	IsActive    bool      `gorm:"column:is_active"`
	SortOrder   int       `gorm:"column:sort_order"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func ToDomainCategory(list []CategoryDb) []models.Category {
	result := make([]models.Category, 0, len(list))
	for i := range list {
		result = append(result, list[i].ToDomainModel())
	}
	return result
}

func (s *CategoryDb) ToDomainModel() models.Category {
	return models.Category{
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

func (s *CategoryDb) Load(m models.Category) {
	s.ID = m.ID
	s.RootID = m.RootID
	s.ParentID = m.ParentID
	s.Name = m.Name
	s.Slug = m.Slug
	s.Description = m.Description
	s.Icon = m.Icon
	s.IsActive = m.IsActive
	s.SortOrder = m.SortOrder
	s.CreatedAt = m.CreatedAt
	s.UpdatedAt = m.UpdatedAt
}

type ParamsCategorySearchDb struct {
	ID        int64
	RootID    int64
	ParentID  int64
	Name      string
	Slug      string
	Icon      string
	IsActive  bool
	SortOrder int
	Limit     int
	Offset    int
}

func (p *ParamsCategorySearchDb) GetQueryRoles() (string, []interface{}) {
	query := []string{}
	params := []interface{}{}

	if p.ID > 0 {
		query = append(query, "id = ? ")
		params = append(params, p.ID)
	}

	if p.RootID > 0 {
		query = append(query, "root_id = ? ")
		params = append(params, p.RootID)
	}

	if p.ParentID > 0 {
		query = append(query, "parent_id = ? ")
		params = append(params, p.ParentID)
	}

	if p.Name > "" {
		query = append(query, "name = ? ")
		params = append(params, p.Name)
	}

	if p.Slug > "" {
		query = append(query, "slug = ? ")
		params = append(params, p.Slug)
	}

	if p.Icon > "" {
		query = append(query, "icon = ? ")
		params = append(params, p.Icon)
	}

	if p.IsActive == true {
		query = append(query, "is_active = ? ")
		params = append(params, p.IsActive)
	}

	if p.SortOrder > 0 {
		query = append(query, "sort_order = ? ")
		params = append(params, p.SortOrder)
	}

	return strings.Join(query, " AND "), params
}

func (p *ParamsCategorySearchDb) ToDB(u *models.ParamsCategorySearch) {
	p.ID = u.ID
	p.RootID = u.RootID
	p.ParentID = u.ParentID
	p.Name = u.Name
	p.Slug = u.Slug
	p.Icon = u.Icon
	p.IsActive = u.IsActive
	p.SortOrder = u.SortOrder
	p.Limit = u.Limit
	p.Offset = u.Offset
}

func (c *CategoryDb) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Local()
	if c.CreatedAt.IsZero() {
		c.CreatedAt = now
	}
	c.UpdatedAt = now
	return nil
}

func (c *CategoryDb) BeforeSave(tx *gorm.DB) (err error) {
	now := time.Now().Local()
	if c.CreatedAt.IsZero() {
		c.CreatedAt = now
	}
	c.UpdatedAt = now
	return nil
}
