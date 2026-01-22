package repositories_test

import (
	"github.com/stretchr/testify/assert"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/repositories"
	modelsDb "taska-core-me-go/cmd/api/repositories/models"
	"taska-core-me-go/cmd/api/utils"
	"taska-core-me-go/cmd/api/utils/data_bases"
	"testing"
)

func Test_FindBy(t *testing.T) {
	assert := assert.New(t)

	conn, err := data_bases.GetTestConnection()
	assert.NoError(err)

	repository := &repositories.SkillsRepository{Conn: conn}
	header := utils.GetTestRequestWithHeaders()
	ctx := utils.CreateRequest(header)

	request := GetParamsSkillsSearch()
	t.Run("Ok", func(t *testing.T) {
		data_bases.CreateTable(ctx, conn, modelsDb.SkillsDb{})

		skillsDb := modelsDb.SkillsDb{
			ID:   1,
			Name: "test",
		}
		repository.Conn.Create(&skillsDb)

		result, err := repository.FindBy(ctx, request)
		assert.NoError(err)
		assert.Equal(1, len(result))
		data_bases.DropTable(ctx, conn, []modelsDb.SkillsDb{})
	})

	t.Run("Error", func(t *testing.T) {
		data_bases.DropTable(ctx, conn, []modelsDb.SkillsDb{})

		_, err = repository.FindBy(ctx, request)
		assert.Error(err)
		data_bases.DropTable(ctx, conn, modelsDb.SkillsDb{})
	})

}

func Test_FindAll(t *testing.T) {
	assert := assert.New(t)

	conn, err := data_bases.GetTestConnection()
	assert.NoError(err)

	repository := &repositories.SkillsRepository{Conn: conn}
	header := utils.GetTestRequestWithHeaders()
	ctx := utils.CreateRequest(header)
	
	t.Run("Ok", func(t *testing.T) {
		data_bases.CreateTable(ctx, conn, modelsDb.SkillsDb{})

		skillsDb := modelsDb.SkillsDb{
			ID:   1,
			Name: "test",
		}
		repository.Conn.Create(&skillsDb)

		result, err := repository.FindAll(ctx)
		assert.NoError(err)
		assert.Equal(1, len(result))
		data_bases.DropTable(ctx, conn, []modelsDb.SkillsDb{})
	})

	t.Run("Error", func(t *testing.T) {
		data_bases.DropTable(ctx, conn, []modelsDb.SkillsDb{})

		_, err = repository.FindAll(ctx)
		assert.Error(err)
		data_bases.DropTable(ctx, conn, modelsDb.SkillsDb{})
	})

}

func GetParamsSkillsSearch() models.ParamsSkillsSearch {
	return models.ParamsSkillsSearch{
		ID:    1,
		Limit: 1,
	}
}
