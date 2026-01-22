package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/response_capture"
	"gorm.io/gorm"
	"net/http"
	"taska-core-me-go/cmd/api/constants"
	"taska-core-me-go/cmd/api/models"
	modelsDB "taska-core-me-go/cmd/api/repositories/models"
)

//go:generate mockgen -destination=../mocks/repositories/$GOFILE -package=mrepositories -source=./$GOFILE

type ISkillsRepository interface {
	FindBy(ctx context.Context, request models.ParamsSkillsSearch) ([]models.SkillsResponse, error)
	FindAll(ctx context.Context) ([]models.SkillsResponse, error)
}

const skillsTableName = "skills"

type SkillsRepository struct {
	Conn *gorm.DB
}

func (repository *SkillsRepository) FindBy(ctx context.Context, request models.ParamsSkillsSearch) ([]models.SkillsResponse, error) {
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFindAll, fmt.Sprintf("Buscando skill con filtro: %+v", request))

	var skillsDb []modelsDB.SkillsDb
	var paramUserDB modelsDB.ParamsSkillsSearchDb
	paramUserDB.ToDB(&request)

	query, params := paramUserDB.GetQueryRoles()

	res := repository.Conn.WithContext(ctx).
		Debug().
		WithContext(ctx).
		Table(skillsTableName).
		Where(query, params...).
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&skillsDb)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			logger.StandardError(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFirstBy, res.Error, "No se encontraron skills con los filtros especificados")
			return nil, response_capture.NewErrorME(ctx, http.StatusNotFound, res.Error, constants.ErrorMessageUserNotFoundByUserName)
		}

		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFirstBy, res.Error, "Error al ejecutar la consulta de skills")
		return nil, response_capture.NewErrorME(ctx, http.StatusBadRequest, res.Error, constants.ErrorMessageErrorFindingUser)
	}

	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFirstBy, fmt.Sprintf("Skills encontrado: %+v", skillsDb))
	return modelsDB.ToDomainList(skillsDb), nil
}

func (repository *SkillsRepository) FindAll(ctx context.Context) ([]models.SkillsResponse, error) {
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFindAll, "Iniciando consulta de todos los skills")
	var skillsListDb []modelsDB.SkillsDb

	res := repository.Conn.
		Debug().
		WithContext(ctx).
		Table(skillsTableName).
		Find(&skillsListDb)

	if res.Error != nil {
		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFindAll, res.Error, "Error al consultar los skills")
		return []models.SkillsResponse{}, response_capture.NewErrorME(ctx, http.StatusBadRequest, res.Error, constants.ErrorMessageErrorFindingUser)
	}

	return modelsDB.ToDomainList(skillsListDb), nil
}
