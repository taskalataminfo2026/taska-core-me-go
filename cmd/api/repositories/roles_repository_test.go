package repositories_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	models2 "github.com/taskalataminfo2026/taska-auth-me-go/cmd/api/models"
	"github.com/taskalataminfo2026/taska-auth-me-go/cmd/api/repositories"
	"github.com/taskalataminfo2026/taska-auth-me-go/cmd/api/repositories/models"
	"github.com/taskalataminfo2026/taska-auth-me-go/cmd/api/utils"
)

func TestRolesRepository_FindAll(t *testing.T) {
	assert := assert.New(t)

	// 🔧 Preparar conexión y repositorio
	conn, err := utils.GetTestConnection()
	assert.NoError(err)

	repository := &repositories.RolesRepository{Conn: conn}
	header := utils.GetTestRequestWithHeaders()
	ctx := utils.CreateRequest(header)

	t.Run("Ok", func(t *testing.T) {
		utils.DropTable(ctx, conn, models.RoleDb{})
		utils.CreateTable(ctx, conn, models.RoleDb{})

		roleDb := models.RoleDb{
			ID:    1,
			Name:  "Admin",
			Level: 1,
		}
		repository.Conn.Create(&roleDb)

		roles, err := repository.FindAll(ctx)
		assert.NoError(err)
		assert.Len(roles, 1)
		assert.Equal("Admin", roles[0].Name)

		utils.DropTable(ctx, conn, models.RoleDb{})
	})

	t.Run("Error", func(t *testing.T) {
		utils.DropTable(ctx, conn, models.RoleDb{})

		roles, err := repository.FindAll(ctx)
		assert.Error(err)
		assert.Nil(roles)
	})

	t.Run("EmptyTable", func(t *testing.T) {
		utils.DropTable(ctx, conn, models.RoleDb{})
		utils.CreateTable(ctx, conn, models.RoleDb{})

		roles, err := repository.FindAll(ctx)
		assert.NoError(err)
		assert.Empty(roles)

		utils.DropTable(ctx, conn, models.RoleDb{})
	})
}

func TestRolesRepository_FirstBy(t *testing.T) {
	assert := assert.New(t)

	conn, err := utils.GetTestConnection()
	assert.NoError(err)

	repository := &repositories.RolesRepository{Conn: conn}
	header := utils.GetTestRequestWithHeaders()
	ctx := utils.CreateRequest(header)

	param := models2.ParamRole{
		ID:    1,
		Name:  "Admin",
		Level: 1,
	}

	t.Run("Ok", func(t *testing.T) {
		utils.DropTable(ctx, conn, models.RoleDb{})
		utils.CreateTable(ctx, conn, models.RoleDb{})

		roleDb := models.RoleDb{
			ID:    1,
			Name:  "Admin",
			Level: 1,
		}
		repository.Conn.Create(&roleDb)

		roles, err := repository.FirstBy(ctx, param)
		assert.NoError(err)
		assert.NotEmpty(roles)
		assert.Equal("Admin", roles[0].Name)

		utils.DropTable(ctx, conn, models.RoleDb{})
	})

	t.Run("ErrRecordNotFound", func(t *testing.T) {
		utils.DropTable(ctx, conn, models.RoleDb{})
		utils.CreateTable(ctx, conn, models.RoleDb{})

		roles, err := repository.FirstBy(ctx, param)
		assert.Error(err)
		assert.Nil(roles)

		utils.DropTable(ctx, conn, models.RoleDb{})
	})

	t.Run("Error_DB", func(t *testing.T) {
		utils.DropTable(ctx, conn, models.RoleDb{})

		roles, err := repository.FirstBy(ctx, param)
		assert.Error(err)
		assert.Nil(roles)
	})
}
