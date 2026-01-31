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

func Test_Category_FindBy(t *testing.T) {
	assert := assert.New(t)

	conn, err := data_bases.GetTestConnection()
	assert.NoError(err)

	repository := &repositories.CategoriesRepository{Conn: conn}
	header := utils.GetTestRequestWithHeaders()
	ctx := utils.CreateRequest(header)

	request := GetParamsCategorySearch()
	t.Run("Ok", func(t *testing.T) {
		data_bases.CreateTable(ctx, conn, modelsDb.CategoryDb{})

		categoryDb := modelsDb.CategoryDb{
			ID:          1,
			Name:        "test category",
			Slug:        "test-category",
			Description: "test description",
			IsActive:    true,
		}
		repository.Conn.Create(&categoryDb)

		result, err := repository.FindBy(ctx, request)
		assert.NoError(err)
		assert.Equal(1, len(result))
		data_bases.DropTable(ctx, conn, []modelsDb.CategoryDb{})
	})

	t.Run("Error", func(t *testing.T) {
		data_bases.DropTable(ctx, conn, []modelsDb.CategoryDb{})

		_, err = repository.FindBy(ctx, request)
		assert.Error(err)
		data_bases.DropTable(ctx, conn, modelsDb.CategoryDb{})
	})

}

func Test_Category_FindAll(t *testing.T) {
	assert := assert.New(t)

	conn, err := data_bases.GetTestConnection()
	assert.NoError(err)

	repository := &repositories.CategoriesRepository{Conn: conn}
	header := utils.GetTestRequestWithHeaders()
	ctx := utils.CreateRequest(header)

	t.Run("Ok", func(t *testing.T) {
		data_bases.CreateTable(ctx, conn, modelsDb.CategoryDb{})

		categoryDb := modelsDb.CategoryDb{
			ID:          1,
			Name:        "test category",
			Slug:        "test-category",
			Description: "test description",
			IsActive:    true,
		}
		repository.Conn.Create(&categoryDb)

		result, err := repository.FindAll(ctx)
		assert.NoError(err)
		assert.Equal(1, len(result))
		data_bases.DropTable(ctx, conn, []modelsDb.CategoryDb{})
	})

	t.Run("Error", func(t *testing.T) {
		data_bases.DropTable(ctx, conn, []modelsDb.CategoryDb{})

		_, err = repository.FindAll(ctx)
		assert.Error(err)
		data_bases.DropTable(ctx, conn, modelsDb.CategoryDb{})
	})

}

func Test_Category_FirstBy(t *testing.T) {
	assert := assert.New(t)

	conn, err := data_bases.GetTestConnection()
	assert.NoError(err)

	repository := &repositories.CategoriesRepository{Conn: conn}
	header := utils.GetTestRequestWithHeaders()
	ctx := utils.CreateRequest(header)

	request := GetParamsCategorySearch()
	t.Run("Ok", func(t *testing.T) {
		data_bases.CreateTable(ctx, conn, modelsDb.CategoryDb{})

		categoryDb := modelsDb.CategoryDb{
			ID:          1,
			Name:        "test category",
			Slug:        "test-category",
			Description: "test description",
			IsActive:    true,
		}
		repository.Conn.Create(&categoryDb)

		result, err := repository.FirstBy(ctx, request)
		assert.NoError(err)
		assert.Equal(int64(1), result.ID)
		data_bases.DropTable(ctx, conn, modelsDb.CategoryDb{})
	})

	t.Run("Error", func(t *testing.T) {
		data_bases.DropTable(ctx, conn, modelsDb.CategoryDb{})

		_, err = repository.FirstBy(ctx, request)
		assert.Error(err)
		data_bases.DropTable(ctx, conn, modelsDb.CategoryDb{})
	})

	t.Run("ErrRecordNotFound", func(t *testing.T) {
		data_bases.DropTable(ctx, conn, modelsDb.CategoryDb{})
		data_bases.CreateTable(ctx, conn, modelsDb.CategoryDb{})

		category, err := repository.FirstBy(ctx, request)
		assert.Error(err)
		assert.Equal(int64(0), category.ID)
		data_bases.DropTable(ctx, conn, modelsDb.CategoryDb{})
	})

}

func Test_Category_Upsert(t *testing.T) {
	assert := assert.New(t)

	conn, err := data_bases.GetTestConnection()
	assert.NoError(err)

	repository := &repositories.CategoriesRepository{Conn: conn}
	header := utils.GetTestRequestWithHeaders()
	ctx := utils.CreateRequest(header)

	request := GetCategoryUpsert()
	t.Run("Save_Ok", func(t *testing.T) {
		data_bases.CreateTable(ctx, conn, modelsDb.CategoryDb{})

		categoryDb := modelsDb.CategoryDb{
			ID:          1,
			Name:        "test category",
			Slug:        "test-category",
			Description: "test description",
			IsActive:    true,
		}
		repository.Conn.Create(&categoryDb)

		result, err := repository.Upsert(ctx, request)
		assert.NoError(err)
		assert.Equal(int64(1), result.ID)
		data_bases.DropTable(ctx, conn, []modelsDb.CategoryDb{})
	})

	t.Run("Create_Ok", func(t *testing.T) {
		data_bases.CreateTable(ctx, conn, modelsDb.CategoryDb{})

		createRequest := models.Category{
			Name:        "new category",
			Slug:        "new-category",
			Description: "new description",
			IsActive:    true,
		}

		result, err := repository.Upsert(ctx, createRequest)
		assert.NoError(err)
		assert.Greater(result.ID, int64(0))
		data_bases.DropTable(ctx, conn, []modelsDb.CategoryDb{})
	})

	t.Run("Update_Ok", func(t *testing.T) {
		data_bases.CreateTable(ctx, conn, modelsDb.CategoryDb{})

		categoryDb := modelsDb.CategoryDb{}
		repository.Conn.Create(&categoryDb)

		result, err := repository.Upsert(ctx, request)
		assert.NoError(err)
		assert.Equal(int64(1), result.ID)
		data_bases.DropTable(ctx, conn, []modelsDb.CategoryDb{})
	})

	t.Run("Error", func(t *testing.T) {
		data_bases.DropTable(ctx, conn, []modelsDb.CategoryDb{})

		_, err = repository.Upsert(ctx, request)
		assert.Error(err)
		data_bases.DropTable(ctx, conn, modelsDb.CategoryDb{})
	})

}

func GetParamsCategorySearch() models.ParamsCategorySearch {
	return models.ParamsCategorySearch{
		ID:    1,
		Limit: 1,
	}
}

func GetCategoryUpsert() models.Category {
	return models.Category{
		ID:          1,
		Name:        "updated category",
		Slug:        "updated-category",
		Description: "updated description",
		IsActive:    true,
	}
}
