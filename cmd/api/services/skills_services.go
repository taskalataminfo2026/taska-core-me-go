package services

import (
	"context"
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
	return services.SkillsRepository.FindBy(ctx, request)
}

func (services *SkillsServices) List(ctx context.Context) ([]models.Skills, error) {
	return services.SkillsRepository.FindAll(ctx)
}

func (services *SkillsServices) Save(ctx context.Context, request models.ParamsSkillsSave) (models.Skills, error) {

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

	skills := models.Skills{
		ID:                   id,
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
