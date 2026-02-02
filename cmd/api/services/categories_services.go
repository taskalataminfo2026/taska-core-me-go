package services

import (
	"context"
	"fmt"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
	"taska-core-me-go/cmd/api/constants"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/repositories"
)

//go:generate mockgen -destination=../mocks/services/$GOFILE -package=mservices -source=./$GOFILE

type ICategoriesServices interface {
	Search(ctx context.Context, request models.ParamsCategorySearch) ([]models.Category, error)
	List(ctx context.Context) ([]models.Category, error)
	Save(ctx context.Context, request models.ParamsCategorySave) (models.Category, error)
	Update(ctx context.Context, id int64, request models.ParamsCategorySave) (models.Category, error)
}

type CategoriesServices struct {
	CategoriesRepository repositories.ICategoriesRepository
}

func (services *CategoriesServices) Search(ctx context.Context, request models.ParamsCategorySearch) ([]models.Category, error) {
	logger.StandardInfo(ctx, constants.LayerService, constants.ModuleCategories, constants.FunctionCategorySearch, fmt.Sprintf("Iniciando búsqueda de categorías con criterios: %v", request))
	return services.CategoriesRepository.FindBy(ctx, request)
}

func (services *CategoriesServices) List(ctx context.Context) ([]models.Category, error) {
	logger.StandardInfo(ctx, constants.LayerService, constants.ModuleCategories, constants.FunctionCategoryList, "Iniciando listado de categorías")
	return services.CategoriesRepository.FindAll(ctx)
}

func (services *CategoriesServices) Save(ctx context.Context, request models.ParamsCategorySave) (models.Category, error) {
	logger.StandardInfo(ctx, constants.LayerService, constants.ModuleCategories, constants.FunctionCategoriesSave, fmt.Sprintf("Creando nueva categoría: name=%s", request.Name))
	category := models.Category{
		RootID:      request.RootID,
		ParentID:    request.ParentID,
		Name:        request.Name,
		Slug:        request.Slug,
		Description: request.Description,
		Icon:        request.Icon,
		IsActive:    request.IsActive,
		SortOrder:   request.SortOrder,
	}

	return services.CategoriesRepository.Upsert(ctx, category)
}

func (services *CategoriesServices) Update(ctx context.Context, id int64, request models.ParamsCategorySave) (models.Category, error) {
	logger.StandardInfo(ctx, constants.LayerService, constants.ModuleCategories, constants.FunctionCategoriesUpsert, fmt.Sprintf("Iniciando actualización de categoría. ID=%d", id))
	var (
		category models.Category
		err      error
	)

	category, err = services.CategoriesRepository.FirstBy(ctx, models.ParamsCategorySearch{ID: id})
	if err != nil {
		return models.Category{}, err
	}

	category.RootID = request.RootID
	category.ParentID = request.ParentID
	category.Name = request.Name
	category.Slug = request.Slug
	category.Description = request.Description
	category.Icon = request.Icon
	category.IsActive = request.IsActive
	category.SortOrder = request.SortOrder

	return services.CategoriesRepository.Upsert(ctx, category)
}
