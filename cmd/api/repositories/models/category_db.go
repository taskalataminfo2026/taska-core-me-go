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
	ID                   int64
	Slug                 string
	AvgPriceEstimate     float64
	RequiresVerification bool
	RiskLevel            int64
	IsActive             bool
	Limit                int
	Offset               int
}

func (p *ParamsCategorySearchDb) GetQueryRoles() (string, []interface{}) {
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

func (p *ParamsCategorySearchDb) ToDB(u *models.ParamsCategorysSearch) {
	p.ID = u.ID
	p.Slug = u.Slug
	p.AvgPriceEstimate = u.AvgPriceEstimate
	p.RequiresVerification = u.RequiresVerification
	p.RiskLevel = u.RiskLevel
	p.IsActive = u.IsActive
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
