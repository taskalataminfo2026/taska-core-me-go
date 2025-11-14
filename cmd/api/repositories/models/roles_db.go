package models

import (
	"gorm.io/gorm"
	"strings"
	"taska-core-me-go/cmd/api/models"
	"time"
)

func (RoleDb) TableName() string { return "roles" }

type ListRoleDb struct {
	RoleDb RoleDb
}

type RoleDb struct {
	ID          int64     `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	Level       int       `gorm:"column:level"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"type:timestamp" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"type:timestamp" json:"updated_at,omitempty"`
}

func ToDomainList(rolesDb []RoleDb) []models.Role {
	roles := make([]models.Role, len(rolesDb))
	for i, role := range rolesDb {
		roles[i] = role.ToDomainModel()
	}
	return roles
}

func (u *RoleDb) ToDomainModel() models.Role {
	return models.Role{
		ID:          u.ID,
		Name:        u.Name,
		Level:       u.Level,
		Description: u.Description,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

type ParamRoleDb struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

func (u *ParamRoleDb) GetQueryRoles() (string, []interface{}) {
	query := []string{}
	params := []interface{}{}

	if u.ID > 0 {
		query = append(query, "id = ? ")
		params = append(params, u.ID)
	}

	if u.Name > "" {
		query = append(query, "name = ? ")
		params = append(params, u.Name)
	}

	if u.Level > 0 {
		query = append(query, "level = ? ")
		params = append(params, u.Level)
	}

	return strings.Join(query, " AND "), params
}

func (db *ParamRoleDb) ToDB(u *models.ParamRole) {
	db.ID = u.ID
	db.Name = u.Name
	db.Level = u.Level
}

func (u *RoleDb) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Local()
	u.CreatedAt = now
	u.UpdatedAt = now
	return nil
}
