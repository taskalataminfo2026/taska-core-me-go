package models

import (
	"time"
)

type ListRole struct {
	RoleDb Role
}

type Role struct {
	ID          int64
	Name        string
	Level       int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ParamRole struct {
	ID    int64
	Name  string
	Level int
}
