package repositories_test

import (
	"github.com/stretchr/testify/assert"
	"taska-core-me-go/cmd/api/utils/data_bases"
	"testing"

	models2 "taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/repositories"
	"taska-core-me-go/cmd/api/repositories/models"
	"taska-core-me-go/cmd/api/utils"
)

func TestRolesRepository_FindAll(t *testing.T) {
	assert := assert.New(t)

	// 🔧 Preparar conexión y repositorio
	conn, err := data_bases.GetTestConnection()
	assert.NoError(err)

	repository := &repositories.RolesRepository{Conn: conn}
	header := utils.GetTestRequestWithHeaders()
	ctx := utils.CreateRequest(header)

	t.Run("Ok", func(t *testing.T) {
		data_bases.DropTable(ctx, conn, models.RoleDb{})
		data_bases.CreateTable(ctx, conn, models.RoleDb{})

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

		data_bases.DropTable(ctx, conn, models.RoleDb{})
	})

	t.Run("Error", func(t *testing.T) {
		data_bases.DropTable(ctx, conn, models.RoleDb{})

		roles, err := repository.FindAll(ctx)
		assert.Error(err)
		assert.Nil(roles)
	})

	t.Run("EmptyTable", func(t *testing.T) {
		data_bases.DropTable(ctx, conn, models.RoleDb{})
		data_bases.CreateTable(ctx, conn, models.RoleDb{})

		roles, err := repository.FindAll(ctx)
		assert.NoError(err)
		assert.Empty(roles)

		data_bases.DropTable(ctx, conn, models.RoleDb{})
	})
}

func TestRolesRepository_FirstBy(t *testing.T) {
	assert := assert.New(t)

	conn, err := data_bases.GetTestConnection()
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
		data_bases.DropTable(ctx, conn, models.RoleDb{})
		data_bases.CreateTable(ctx, conn, models.RoleDb{})

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

		data_bases.DropTable(ctx, conn, models.RoleDb{})
	})

	t.Run("ErrRecordNotFound", func(t *testing.T) {
		data_bases.DropTable(ctx, conn, models.RoleDb{})
		data_bases.CreateTable(ctx, conn, models.RoleDb{})

		roles, err := repository.FirstBy(ctx, param)
		assert.Error(err)
		assert.Nil(roles)

		data_bases.DropTable(ctx, conn, models.RoleDb{})
	})

	t.Run("Error_DB", func(t *testing.T) {
		data_bases.DropTable(ctx, conn, models.RoleDb{})

		roles, err := repository.FirstBy(ctx, param)
		assert.Error(err)
		assert.Nil(roles)
	})
}
