package providers

import (
	"gorm.io/gorm"
	"taska-core-me-go/cmd/api/repositories"
)

func CategoriesRepository(conn *gorm.DB) *repositories.CategoriesRepository {
	return &repositories.CategoriesRepository{Conn: conn}
}

<<<<<<< HEAD
func RolesTokenRepository(conn *gorm.DB) *repositories.RolesRepository {
	return &repositories.RolesRepository{Conn: conn}
}

func BlacklistedTokenRepository(conn *gorm.DB) *repositories.BlacklistedTokenRepository {
	return &repositories.BlacklistedTokenRepository{Conn: conn}
}

func SkillsCategoriesRepository(conn *gorm.DB) *repositories.SkillsCategoriesRepository {
	return &repositories.SkillsCategoriesRepository{Conn: conn}
}

=======
>>>>>>> master
func SkillsRepository(conn *gorm.DB) *repositories.SkillsRepository {
	return &repositories.SkillsRepository{Conn: conn}
}
