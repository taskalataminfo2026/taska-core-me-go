package repositories_test

import (
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/repositories"
	modelsDb "taska-core-me-go/cmd/api/repositories/models"
	"taska-core-me-go/cmd/api/utils"
	"taska-core-me-go/cmd/api/utils/data_bases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SkillCategory_FirstBy(t *testing.T) {
	assert := assert.New(t)

	conn, err := data_bases.GetTestConnection()
	assert.NoError(err)

	repository := &repositories.SkillsCategoriesRepository{Conn: conn}
	header := utils.GetTestRequestWithHeaders()
	ctx := utils.CreateRequest(header)

	request := GetParamsSkillsCategorySearch()
	t.Run("Ok", func(t *testing.T) {
		data_bases.CreateTable(ctx, conn, modelsDb.SkillCategoryDb{})

		skillCategoryDb := modelsDb.SkillCategoryDb{
			ID:         1,
			SkillID:    1,
			CategoryID: 1,
			IsPrimary:  true,
			IsActive:   true,
		}
		repository.Conn.Create(&skillCategoryDb)

		result, err := repository.FirstBy(ctx, request)
		assert.NoError(err)
		assert.Equal(int64(1), result.ID)
		data_bases.DropTable(ctx, conn, modelsDb.SkillCategoryDb{})
	})

	t.Run("Error", func(t *testing.T) {
		data_bases.DropTable(ctx, conn, modelsDb.SkillCategoryDb{})

		_, err = repository.FirstBy(ctx, request)
		assert.Error(err)
		data_bases.DropTable(ctx, conn, modelsDb.SkillCategoryDb{})
	})

	t.Run("ErrRecordNotFound", func(t *testing.T) {
		data_bases.DropTable(ctx, conn, modelsDb.SkillCategoryDb{})
		data_bases.CreateTable(ctx, conn, modelsDb.SkillCategoryDb{})

		skillCategory, err := repository.FirstBy(ctx, request)
		assert.Error(err)
		assert.Equal(int64(0), skillCategory.ID)
		data_bases.DropTable(ctx, conn, modelsDb.SkillCategoryDb{})
	})

}

func Test_SkillCategory_Upsert(t *testing.T) {
	assert := assert.New(t)

	conn, err := data_bases.GetTestConnection()
	assert.NoError(err)

	repository := &repositories.SkillsCategoriesRepository{Conn: conn}
	header := utils.GetTestRequestWithHeaders()
	ctx := utils.CreateRequest(header)

	request := GetSkillCategoryUpsert()
	t.Run("Save_Ok", func(t *testing.T) {
		data_bases.CreateTable(ctx, conn, modelsDb.SkillCategoryDb{})

		skillCategoryDb := modelsDb.SkillCategoryDb{
			ID:         1,
			SkillID:    1,
			CategoryID: 1,
			IsPrimary:  true,
			IsActive:   true,
		}
		repository.Conn.Create(&skillCategoryDb)

		result, err := repository.Upsert(ctx, request)
		assert.NoError(err)
		assert.Equal(int64(1), result.ID)
		data_bases.DropTable(ctx, conn, []modelsDb.SkillCategoryDb{})
	})

	t.Run("Create_Ok", func(t *testing.T) {
		data_bases.CreateTable(ctx, conn, modelsDb.SkillCategoryDb{})

		createRequest := models.SkillCategory{
			SkillID:    2,
			CategoryID: 2,
			IsPrimary:  false,
			IsActive:   true,
		}

		result, err := repository.Upsert(ctx, createRequest)
		assert.NoError(err)
		assert.Greater(result.ID, int64(0))
		data_bases.DropTable(ctx, conn, []modelsDb.SkillCategoryDb{})
	})

	t.Run("Update_Ok", func(t *testing.T) {
		data_bases.CreateTable(ctx, conn, modelsDb.SkillCategoryDb{})

		skillCategoryDb := modelsDb.SkillCategoryDb{}
		repository.Conn.Create(&skillCategoryDb)

		result, err := repository.Upsert(ctx, request)
		assert.NoError(err)
		assert.Equal(int64(1), result.ID)
		data_bases.DropTable(ctx, conn, []modelsDb.SkillCategoryDb{})
	})

	t.Run("Error", func(t *testing.T) {
		data_bases.DropTable(ctx, conn, []modelsDb.SkillCategoryDb{})

		_, err = repository.Upsert(ctx, request)
		assert.Error(err)
		data_bases.DropTable(ctx, conn, modelsDb.SkillCategoryDb{})
	})

}

func GetParamsSkillsCategorySearch() models.ParamsSkillsCategorySearch {
	return models.ParamsSkillsCategorySearch{
		ID:    1,
		Limit: 1,
	}
}

func GetSkillCategoryUpsert() models.SkillCategory {
	return models.SkillCategory{
		ID:         1,
		SkillID:    1,
		CategoryID: 1,
		IsPrimary:  true,
		IsActive:   true,
	}
}
