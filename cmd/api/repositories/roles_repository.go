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

type IRolesRepository interface {
	FindAll(ctx context.Context) ([]models.Role, error)
	FirstBy(ctx context.Context, users models.ParamRole) ([]models.Role, error)
}

const rolesTableName = "roles"

type RolesRepository struct {
	Conn *gorm.DB
}

func (repository *RolesRepository) FindAll(ctx context.Context) ([]models.Role, error) {
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFindAll, "Iniciando consulta de todos los roles")
	var rolesDb []modelsDB.RoleDb

	res := repository.Conn.
		WithContext(ctx).
		Table(rolesTableName).
		Find(&rolesDb)

	if res.Error != nil {
		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFindAll, res.Error, "Error al consultar los roles")
		return nil, response_capture.NewErrorME(ctx, http.StatusBadRequest, res.Error, constants.ErrorMessageErrorFindingUser)
	}

	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFindAll, fmt.Sprintf("Consulta exitosa: %d roles encontrados", len(rolesDb)))
	return modelsDB.ToDomainRoles(rolesDb), nil
}

func (repository *RolesRepository) FirstBy(ctx context.Context, filter models.ParamRole) ([]models.Role, error) {
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFindAll, fmt.Sprintf("Buscando roles con filtro: %+v", filter))

	var rolesDb []modelsDB.RoleDb
	var paramUserDB modelsDB.ParamRoleDb
	paramUserDB.ToDB(&filter)

	query, params := paramUserDB.GetQueryRoles()

	res := repository.Conn.WithContext(ctx).
		Table(rolesTableName).
		Where(query, params...).
		First(&rolesDb)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			logger.StandardError(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFirstBy, res.Error, "No se encontraron roles con los filtros especificados")
			return nil, response_capture.NewErrorME(ctx, http.StatusNotFound, res.Error, constants.ErrorMessageUserNotFoundByUserName)
		}

		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFirstBy, res.Error, "Error al ejecutar la consulta de roles")
		return nil, response_capture.NewErrorME(ctx, http.StatusBadRequest, res.Error, constants.ErrorMessageErrorFindingUser)
	}

	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleRoles, constants.FunctionFirstBy, fmt.Sprintf("Rol encontrado: %+v", rolesDb))
	return modelsDB.ToDomainRoles(rolesDb), nil
}
