package services

import (
	"context"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
	"go.uber.org/zap"
	"taska-core-me-go/cmd/api/constants"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/repositories"
)

//go:generate mockgen -destination=../mocks/services/$GOFILE -package=mservices -source=./$GOFILE

type ISkillsServices interface {
	Search(ctx context.Context, request models.ParamsSkillsSearch) ([]models.Skills, error)
	List(ctx context.Context) ([]models.Skills, error)
	Save(ctx context.Context, request models.ParamsSkillsSave) (models.Skills, error)
	Update(ctx context.Context, id int64, request models.ParamsSkillsSave) (models.Skills, error)
}

type SkillsServices struct {
	SkillsRepository repositories.ISkillsRepository
}

func (services *SkillsServices) Search(ctx context.Context, request models.ParamsSkillsSearch) ([]models.Skills, error) {
	logger.StandardInfo(ctx, constants.LayerService, constants.ModuleSkills, constants.FunctionSkillsSearch, "Búsqueda de skills")
	return services.SkillsRepository.FindBy(ctx, request)
}

func (services *SkillsServices) List(ctx context.Context) ([]models.Skills, error) {
	logger.StandardInfo(ctx, constants.LayerService, constants.ModuleSkills, constants.FunctionSkillsList, "Listado de skills")
	return services.SkillsRepository.FindAll(ctx)
}

func (services *SkillsServices) Save(ctx context.Context, request models.ParamsSkillsSave) (models.Skills, error) {
	logger.StandardInfo(ctx, constants.LayerService, constants.ModuleSkills, constants.FunctionSkillsSave, "Creando skill")
	skills := models.Skills{
		Name:                 request.Name,
		Slug:                 request.Slug,
		Description:          request.Description,
		AvgPriceEstimate:     request.AvgPriceEstimate,
		RequiresVerification: request.RequiresVerification,
		RiskLevel:            request.RiskLevel,
		IsActive:             request.IsActive,
	}

	return services.SkillsRepository.Upsert(ctx, skills)
}

func (services *SkillsServices) Update(ctx context.Context, id int64, request models.ParamsSkillsSave) (models.Skills, error) {
	logger.StandardInfo(ctx, constants.LayerService, constants.ModuleSkills, constants.FunctionSkillsUpsert, "Actualizando skill", zap.Int64("skill_id", id))
	var (
		skills models.Skills
		err    error
	)

	skills, err = services.SkillsRepository.FirstBy(ctx, models.ParamsSkillsSearch{ID: id})
	if err != nil {
		return models.Skills{}, err
	}

	skills.Name = request.Name
	skills.Slug = request.Slug
	skills.Description = request.Description
	skills.AvgPriceEstimate = request.AvgPriceEstimate
	skills.RequiresVerification = request.RequiresVerification
	skills.RiskLevel = request.RiskLevel
	skills.IsActive = request.IsActive

	return services.SkillsRepository.Upsert(ctx, skills)
}
