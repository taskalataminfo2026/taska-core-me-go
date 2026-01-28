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

type ICategoriesRepository interface {
	FindBy(ctx context.Context, request models.ParamsCategorysSearch) ([]models.Category, error)
	FindAll(ctx context.Context) ([]models.Category, error)
	FirstBy(ctx context.Context, request models.ParamsCategorysSearch) (models.Category, error)
	Upsert(ctx context.Context, request models.Category) (models.Category, error)
}

const categoriesTableName = "categories"

type CategoriesRepository struct {
	Conn *gorm.DB
}

func (repository *CategoriesRepository) FindBy(ctx context.Context, request models.ParamsCategorysSearch) ([]models.Category, error) {
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleCategories, constants.FunctionFindAll, fmt.Sprintf("Buscando categoria con filtro: %+v", request))

	var categoryDb []modelsDB.CategoryDb
	var paramUserDB modelsDB.ParamsCategorySearchDb
	paramUserDB.ToDB(&request)

	query, params := paramUserDB.GetQueryRoles()

	res := repository.Conn.WithContext(ctx).
		WithContext(ctx).
		Table(categoriesTableName).
		Where(query, params...).
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&categoryDb)

	if res.Error != nil {
		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleCategories, constants.FunctionFirstBy, res.Error, "Error al ejecutar la consulta de categoria")
		return nil, response_capture.NewErrorME(ctx, http.StatusBadRequest, res.Error, constants.ErrorMessageErrorFindingUser)
	}

	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleCategories, constants.FunctionFirstBy, fmt.Sprintf("Categories encontrado: %+v", categoryDb))
	return modelsDB.ToDomainCategory(categoryDb), nil
}

func (repository *CategoriesRepository) FindAll(ctx context.Context) ([]models.Category, error) {
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleCategories, constants.FunctionFindAll, "Iniciando consulta de todos los categoria")
	var categoryListDb []modelsDB.CategoryDb

	res := repository.Conn.
		Debug().
		WithContext(ctx).
		Table(categoriesTableName).
		Find(&categoryListDb)

	if res.Error != nil {
		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleCategories, constants.FunctionFindAll, res.Error, "Error al consultar los categoria")
		return []models.Category{}, response_capture.NewErrorME(ctx, http.StatusBadRequest, res.Error, constants.ErrorMessageErrorFindingUser)
	}

	return modelsDB.ToDomainCategory(categoryListDb), nil
}

func (repository *CategoriesRepository) FirstBy(ctx context.Context, request models.ParamsCategorysSearch) (models.Category, error) {
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleCategories, constants.FunctionFindAll, fmt.Sprintf("Buscando skill con filtro: %+v", request))

	var skillsDb modelsDB.CategoryDb
	var paramUserDB modelsDB.ParamsCategorySearchDb
	paramUserDB.ToDB(&request)

	query, params := paramUserDB.GetQueryRoles()

	res := repository.Conn.WithContext(ctx).
		WithContext(ctx).
		Table(categoriesTableName).
		Where(query, params...).
		First(&skillsDb)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			logger.StandardError(ctx, constants.LayerRepository, constants.ModuleCategories, constants.FunctionFirstBy, res.Error, "categoria no encontrado por filtros")
			return models.Category{}, response_capture.NewErrorME(ctx, http.StatusNotFound, res.Error, constants.ErrorMessageUserNotFoundByUserName)
		}
		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleCategories, constants.FunctionFirstBy, res.Error, "Error al ejecutar la consulta de categoria")
		return models.Category{}, response_capture.NewErrorME(ctx, http.StatusBadRequest, res.Error, constants.ErrorMessageErrorFindingUser)
	}

	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleCategories, constants.FunctionFirstBy, fmt.Sprintf("Categories encontrado: %+v", skillsDb))
	return skillsDb.ToDomainModel(), nil
}

func (repository *CategoriesRepository) Upsert(ctx context.Context, request models.Category) (models.Category, error) {

	var (
		entity modelsDB.CategoryDb
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
		return entity.ToDomainModel(), response_capture.NewErrorME(ctx, http.StatusBadRequest, err, fmt.Sprintf(constants.ErrorMessageSavingToken, err.Error()))
	}

	tx.Commit()
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleCategories, constants.FunctionUpsert, "Categoria guardado correctamente", zap.Int64("skill_id", entity.ID))
	return entity.ToDomainModel(), nil
}
