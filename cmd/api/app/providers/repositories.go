package providers

import (
	"gorm.io/gorm"
	"taska-core-me-go/cmd/api/repositories"
)

func CategoriesRepository(conn *gorm.DB) *repositories.CategoriesRepository {
	return &repositories.CategoriesRepository{Conn: conn}
}

func SkillsRepository(conn *gorm.DB) *repositories.SkillsRepository {
	return &repositories.SkillsRepository{Conn: conn}
}
