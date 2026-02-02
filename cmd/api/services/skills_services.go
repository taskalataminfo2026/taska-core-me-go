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
	logger.StandardInfo(ctx, constants.LayerService, constants.ModuleSkills, constants.FunctionSkillsSearch, fmt.Sprintf("Iniciando búsqueda de categorías con criterios: %v", request))
	return services.SkillsRepository.FindBy(ctx, request)
}

func (services *SkillsServices) List(ctx context.Context) ([]models.Skills, error) {
	logger.StandardInfo(ctx, constants.LayerService, constants.ModuleSkills, constants.FunctionSkillsList, "Iniciando listado de categorías")
	return services.SkillsRepository.FindAll(ctx)
}

func (services *SkillsServices) Save(ctx context.Context, request models.ParamsSkillsSave) (models.Skills, error) {
	logger.StandardInfo(ctx, constants.LayerService, constants.ModuleSkills, constants.FunctionSkillsSave, fmt.Sprintf("Creando skill: name=%v", request.Name))
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
	logger.StandardInfo(ctx, constants.LayerService, constants.ModuleSkills, constants.FunctionSkillsUpsert, fmt.Sprintf("Iniciando actualización de relación skill-categoría. ID=%d", id))
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
