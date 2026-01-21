package providers

import (
	"gorm.io/gorm"
	"taska-core-me-go/cmd/api/repositories"
)

func RolesTokenRepository(conn *gorm.DB) *repositories.RolesRepository {
	return &repositories.RolesRepository{Conn: conn}
}

func BlacklistedTokenRepository(conn *gorm.DB) *repositories.BlacklistedTokenRepository {
	return &repositories.BlacklistedTokenRepository{Conn: conn}
}

func SkillsRepository(conn *gorm.DB) *repositories.SkillsRepository {
	return &repositories.SkillsRepository{Conn: conn}
}
