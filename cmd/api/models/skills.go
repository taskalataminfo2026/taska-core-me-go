package models

import (
	"time"
)

type ListSkillsResponse struct {
	Items []SkillsResponse
	Total int
}

type SkillsResponse struct {
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

type SkillsRequest struct {
	ID                   int64   `json:"id"`
	Name                 string  `json:"name"`
	Slug                 string  `json:"slug"`
	AvgPriceEstimate     float64 `json:"avg_price_estimate"`
	RequiresVerification bool    `json:"requires_verification"`
	RiskLevel            int64   `json:"risk_level"`
	IsActive             bool    `json:"is_active"`
}
