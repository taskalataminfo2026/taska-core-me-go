package dto

import "time"

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
}
