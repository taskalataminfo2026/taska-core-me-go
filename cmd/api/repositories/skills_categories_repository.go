package repositories

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"taska-core-me-go/cmd/api/constants"
	"taska-core-me-go/cmd/api/models"
	modelsDB "taska-core-me-go/cmd/api/repositories/models"

	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/response_capture"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/repositories/$GOFILE -package=mrepositories -source=./$GOFILE

type ISkillsCategoriesRepository interface {
	FirstBy(ctx context.Context, request models.ParamsSkillsCategorySearch) (models.SkillCategory, error)
	Upsert(ctx context.Context, request models.SkillCategory) (models.SkillCategory, error)
}

const skillsCategoriesTableName = "skill_categories"

type SkillsCategoriesRepository struct {
	Conn *gorm.DB
}

func (repository *SkillsCategoriesRepository) FirstBy(ctx context.Context, request models.ParamsSkillsCategorySearch) (models.SkillCategory, error) {
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleCategories, constants.FunctionFindAll, fmt.Sprintf("Buscando skill con filtro: %+v", request))

	var entity modelsDB.SkillCategoryDb
	var paramUserDB modelsDB.ParamsSkillsCategorySearchDb
	paramUserDB.ToDB(&request)

	query, params := paramUserDB.GetQueryRoles()

	res := repository.Conn.WithContext(ctx).
		WithContext(ctx).
		Table(skillsCategoriesTableName).
		Where(query, params...).
		First(&entity)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			logger.StandardError(ctx, constants.LayerRepository, constants.ModuleCategories, constants.FunctionFirstBy, res.Error, "categoria no encontrado por filtros")
			return models.SkillCategory{}, response_capture.NewErrorME(ctx, http.StatusNotFound, res.Error, constants.ErrorMessageUserNotFoundByUserName)
		}
		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleCategories, constants.FunctionFirstBy, res.Error, "Error al ejecutar la consulta de categoria")
		return models.SkillCategory{}, response_capture.NewErrorME(ctx, http.StatusBadRequest, res.Error, constants.ErrorMessageErrorFindingUser)
	}

	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleCategories, constants.FunctionFirstBy, fmt.Sprintf("Categories encontrado: %+v", entity))
	return entity.ToDomainModel(), nil
}

func (repository *SkillsCategoriesRepository) Upsert(ctx context.Context, request models.SkillCategory) (models.SkillCategory, error) {

	var (
		entity modelsDB.SkillCategoryDb
		err    error
	)

	entity.Load(request)
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleCategories, constants.FunctionFindAll, fmt.Sprintf("Guardando o actualizando categoria: %+v", request))

	tx := repository.Conn.WithContext(ctx).Begin()

	if entity.ID > 0 {
		err = tx.Save(&entity).Error
	} else {
		err = tx.Create(&entity).Error
	}

	if err != nil {
		tx.Rollback()
		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleCategories, constants.FunctionUpsert, err, "Error guardando la categoria", zap.Any("body", entity))
		return models.SkillCategory{}, response_capture.NewErrorME(ctx, http.StatusBadRequest, err, fmt.Sprintf(constants.ErrorMessageSavingToken, err.Error()))
	}

	tx.Commit()
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleCategories, constants.FunctionUpsert, "Categoria guardado correctamente", zap.Int64("skill_id", entity.ID))
	return entity.ToDomainModel(), nil
}
