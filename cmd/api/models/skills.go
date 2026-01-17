package models

import (
	"time"
)

type ListSkills struct {
	RoleDb Skills
}

type Skills struct {
	ID          int64
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ParamSkills struct {
	ID    int64
	Name  string
	Level int
}
