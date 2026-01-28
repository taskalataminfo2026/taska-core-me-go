package models

import "time"

type Category struct {
	ID          int64
	RootID      int64
	ParentID    int64
	Name        string
	Slug        string
	Description string
	Icon        string
	IsActive    bool
	SortOrder   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ParamsCategorysSearch struct {
	ID                   int64
	Slug                 string
	AvgPriceEstimate     float64
	RequiresVerification bool
	RiskLevel            int64
	IsActive             bool
	Limit                int
	Offset               int
}
