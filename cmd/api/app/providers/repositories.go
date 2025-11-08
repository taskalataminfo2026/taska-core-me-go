package providers

import (
	"github.com/taskalataminfo2026/taska-auth-me-go/cmd/api/repositories"
	"gorm.io/gorm"
)

func RolesTokenRepository(conn *gorm.DB) *repositories.RolesRepository {
	return &repositories.RolesRepository{Conn: conn}
}
