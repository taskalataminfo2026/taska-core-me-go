package repositories

import (
	"context"
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
	FindAll(ctx context.Context) ([]models.SkillsResponse, error)
	//FirstBy(ctx context.Context, users dto.ParamRole) (dto.SkillsResponse, error)
}

const skillsTableName = "skills"

type SkillsRepository struct {
	Conn *gorm.DB
}

func (repository *SkillsRepository) FindAll(ctx context.Context) ([]models.SkillsResponse, error) {
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFindAll, "Iniciando consulta de todos los roles")
	var skillsListDb []modelsDB.SkillsDb

	res := repository.Conn.
		Debug().
		WithContext(ctx).
		Table(skillsTableName).
		Find(&skillsListDb)

	if res.Error != nil {
		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFindAll, res.Error, "Error al consultar los roles")
		return []models.SkillsResponse{}, response_capture.NewErrorME(ctx, http.StatusBadRequest, res.Error, constants.ErrorMessageErrorFindingUser)
	}

	return modelsDB.ToDomainList(skillsListDb), nil
}

//func (repository *SkillsRepository) FirstBy(ctx context.Context, filter dto.ParamRole) (dto.SkillsResponse, error) {
//	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFindAll, fmt.Sprintf("Buscando roles con filtro: %+v", filter))
//
//	var skillsDb []modelsDB.SkillsDb
//	var paramUserDB modelsDB.ParamSkillsDb
//	paramUserDB.ToDB(&filter)
//
//	query, params := paramUserDB.GetQueryRoles()
//
//	res := repository.Conn.WithContext(ctx).
//		Table(skillsTableName).
//		Where(query, params...).
//		First(&skillsDb)
//
//	if res.Error != nil {
//		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
//			logger.StandardError(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFirstBy, res.Error, "No se encontraron roles con los filtros especificados")
//			return nil, response_capture.NewErrorME(ctx, http.StatusNotFound, res.Error, constants.ErrorMessageUserNotFoundByUserName)
//		}
//
//		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFirstBy, res.Error, "Error al ejecutar la consulta de roles")
//		return nil, response_capture.NewErrorME(ctx, http.StatusBadRequest, res.Error, constants.ErrorMessageErrorFindingUser)
//	}
//
//	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFirstBy, fmt.Sprintf("Rol encontrado: %+v", skillsDb))
//	return modelsDB.ToDomainList(skillsDb), nil
//}
