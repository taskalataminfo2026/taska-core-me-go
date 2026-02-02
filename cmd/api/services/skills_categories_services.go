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

type ISkillsCategoriesServices interface {
	Save(ctx context.Context, request models.ParamsSkillsCategorySave) (models.SkillCategory, error)
	Update(ctx context.Context, id int64, request models.ParamsSkillsCategorySave) (models.SkillCategory, error)
}

type SkillsCategoriesServices struct {
	SkillsCategoriesRepository repositories.ISkillsCategoriesRepository
}

func (services *SkillsCategoriesServices) Save(ctx context.Context, request models.ParamsSkillsCategorySave) (models.SkillCategory, error) {
	logger.StandardInfo(ctx, constants.LayerService, constants.ModuleSkillsAndCategories, constants.FunctionSkillsCategoriesSave, fmt.Sprintf("Creando relación skill-categoría: skillID=%d, categoryID=%d", request.SkillID, request.CategoryID))
	skillCategory := models.SkillCategory{
		SkillID:    request.SkillID,
		CategoryID: request.CategoryID,
		IsPrimary:  request.IsPrimary,
		IsActive:   request.IsActive,
	}

	return services.SkillsCategoriesRepository.Upsert(ctx, skillCategory)
}

func (services *SkillsCategoriesServices) Update(ctx context.Context, id int64, request models.ParamsSkillsCategorySave) (models.SkillCategory, error) {
	logger.StandardInfo(ctx, constants.LayerService, constants.ModuleSkillsAndCategories, constants.FunctionSkillsCategoriesUpdate, fmt.Sprintf("Iniciando actualización de relación skill-categoría. ID=%d", id))
	var (
		category models.SkillCategory
		err      error
	)

	category, err = services.SkillsCategoriesRepository.FirstBy(ctx, models.ParamsSkillsCategorySearch{ID: id})
	if err != nil {
		return models.SkillCategory{}, err
	}

	category.CategoryID = request.CategoryID
	category.SkillID = request.SkillID
	category.IsPrimary = request.IsPrimary
	category.IsActive = request.IsActive

	return services.SkillsCategoriesRepository.Upsert(ctx, category)
}
