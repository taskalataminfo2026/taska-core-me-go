package models

import (
	"gorm.io/gorm"
	"strings"
	"taska-core-me-go/cmd/api/models"
	"time"
)

func (SkillsDb) TableName() string { return "skills" }

type ListSkillsDb struct {
	SkillsDb SkillsDb
}

type SkillsDb struct {
	ID          int64     `gorm:"column:id"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"type:timestamp" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"type:timestamp" json:"updated_at,omitempty"`
}

func ToDomainList(skillsDb []SkillsDb) []models.Skills {
	skills := make([]models.Skills, len(skillsDb))
	for i, skill := range skillsDb {
		skills[i] = skill.ToDomainModel()
	}
	return skills
}

func (u *SkillsDb) ToDomainModel() models.Skills {
	return models.Skills{
		ID:          u.ID,
		Name:        u.Name,
		Description: u.Description,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

type ParamSkillsDb struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

func (u *ParamSkillsDb) GetQueryRoles() (string, []interface{}) {
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

	return strings.Join(query, " AND "), params
}

func (db *ParamSkillsDb) ToDB(u *models.ParamRole) {
	db.ID = u.ID
	db.Name = u.Name
	db.Level = u.Level
}

func (u *SkillsDb) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().Local()
	u.CreatedAt = now
	u.UpdatedAt = now
	return nil
}
