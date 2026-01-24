package models

import (
	"time"
)

type ListSkillsResponse struct {
	Items []Skills
	Total int
}

type Skills struct {
	ID                   int64
	Name                 string
	Slug                 string
	Description          string
	AvgPriceEstimate     float64
	RequiresVerification bool
	RiskLevel            int64
	IsActive             bool
	CreatedAt            time.Time
}

type ParamsSkillsSearch struct {
	ID                   int64
	Slug                 string
	AvgPriceEstimate     float64
	RequiresVerification bool
	RiskLevel            int64
	IsActive             bool
	Limit                int
	Offset               int
}
