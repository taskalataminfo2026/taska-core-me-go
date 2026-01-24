package models

import (
	"gorm.io/gorm"
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
}

func (c *CategoryDb) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Local()
	c.CreatedAt = now
	return nil
}
