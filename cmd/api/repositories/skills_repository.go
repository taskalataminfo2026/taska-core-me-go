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

type ISkillsRepository interface {
	FindBy(ctx context.Context, request models.ParamsSkillsSearch) ([]models.Skills, error)
	FindAll(ctx context.Context) ([]models.Skills, error)
	FirstBy(ctx context.Context, request models.ParamsSkillsSearch) (models.Skills, error)
	Upsert(ctx context.Context, request models.Skills) (models.Skills, error)
}

const skillsTableName = "skills"

type SkillsRepository struct {
	Conn *gorm.DB
}

func (repository *SkillsRepository) FindBy(ctx context.Context, request models.ParamsSkillsSearch) ([]models.Skills, error) {
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleSkills, constants.FunctionFindAll, fmt.Sprintf("Buscando skill con filtro: %+v", request))

	var skillsDb []modelsDB.SkillsDb
	var paramUserDB modelsDB.ParamsSkillsSearchDb
	paramUserDB.ToDB(&request)

	query, params := paramUserDB.GetQueryRoles()

	res := repository.Conn.WithContext(ctx).
		WithContext(ctx).
		Table(skillsTableName).
		Where(query, params...).
		Limit(request.Limit).
		Offset(request.Offset).
		Find(&skillsDb)

	if res.Error != nil {
		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleSkills, constants.FunctionFirstBy, res.Error, "Error al ejecutar la consulta de skills")
		return nil, response_capture.NewErrorME(ctx, http.StatusBadRequest, res.Error, constants.ErrorMessageErrorFindingUser)
	}

	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleSkills, constants.FunctionFirstBy, fmt.Sprintf("Skills encontrado: %+v", skillsDb))
	return modelsDB.ToDomainList(skillsDb), nil
}

func (repository *SkillsRepository) FindAll(ctx context.Context) ([]models.Skills, error) {
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleSkills, constants.FunctionFindAll, "Iniciando consulta de todos los skills")
	var skillsListDb []modelsDB.SkillsDb

	res := repository.Conn.
		Debug().
		WithContext(ctx).
		Table(skillsTableName).
		Find(&skillsListDb)

	if res.Error != nil {
		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleSkills, constants.FunctionFindAll, res.Error, "Error al consultar los skills")
		return []models.Skills{}, response_capture.NewErrorME(ctx, http.StatusBadRequest, res.Error, constants.ErrorMessageErrorFindingUser)
	}

	return modelsDB.ToDomainList(skillsListDb), nil
}

func (repository *SkillsRepository) FirstBy(ctx context.Context, request models.ParamsSkillsSearch) (models.Skills, error) {
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleSkills, constants.FunctionFindAll, fmt.Sprintf("Buscando skill con filtro: %+v", request))

	var skillsDb modelsDB.SkillsDb
	var paramUserDB modelsDB.ParamsSkillsSearchDb
	paramUserDB.ToDB(&request)

	query, params := paramUserDB.GetQueryRoles()

	res := repository.Conn.WithContext(ctx).
		WithContext(ctx).
		Table(skillsTableName).
		Where(query, params...).
		First(&skillsDb)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			logger.StandardError(ctx, constants.LayerRepository, constants.ModuleSkills, constants.FunctionFirstBy, res.Error, "Skills no encontrado por filtros")
			return models.Skills{}, response_capture.NewErrorME(ctx, http.StatusNotFound, res.Error, constants.ErrorMessageUserNotFoundByUserName)
		}
		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleSkills, constants.FunctionFirstBy, res.Error, "Error al ejecutar la consulta de skills")
		return models.Skills{}, response_capture.NewErrorME(ctx, http.StatusBadRequest, res.Error, constants.ErrorMessageErrorFindingUser)
	}

	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleSkills, constants.FunctionFirstBy, fmt.Sprintf("Skills encontrado: %+v", skillsDb))
	return skillsDb.ToDomainModel(), nil
}

func (repository *SkillsRepository) Upsert(ctx context.Context, request models.Skills) (models.Skills, error) {

	var (
		entity modelsDB.SkillsDb
		err    error
	)

	entity.Load(request)
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleSkills, constants.FunctionFindAll, fmt.Sprintf("Guardando o actualizando skill: %+v", request))

	tx := repository.Conn.WithContext(ctx).Begin()

	if entity.ID > 0 {
		err = tx.Save(&entity).Error
	} else {
		err = tx.Create(&entity).Error
	}

	if err != nil {
		tx.Rollback()
		logger.StandardError(ctx, constants.LayerRepository, constants.ModuleSkills, constants.FunctionUpsert, err, "Error guardando el skill", zap.Any("body", entity))
		return entity.ToDomainModel(), response_capture.NewErrorME(ctx, http.StatusBadRequest, err, fmt.Sprintf(constants.ErrorMessageSavingToken, err.Error()))
	}

	tx.Commit()
	logger.StandardInfo(ctx, constants.LayerRepository, constants.ModuleSkills, constants.FunctionUpsert, "skill guardado correctamente", zap.Int64("skill_id", entity.ID))
	return entity.ToDomainModel(), nil
}
