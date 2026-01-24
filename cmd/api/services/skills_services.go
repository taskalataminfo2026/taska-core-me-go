package services

import (
	"context"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/repositories"
)

//go:generate mockgen -destination=../mocks/services/$GOFILE -package=mservices -source=./$GOFILE

type ISkillsServices interface {
	SkillsSearch(ctx context.Context, request models.ParamsSkillsSearch) ([]models.Skills, error)
	SkillsList(ctx context.Context) ([]models.Skills, error)
}

type SkillsServices struct {
	SkillsRepository repositories.ISkillsRepository
}

func (services *SkillsServices) SkillsSearch(ctx context.Context, request models.ParamsSkillsSearch) ([]models.Skills, error) {
	return services.SkillsRepository.FindBy(ctx, request)
}

func (services *SkillsServices) SkillsList(ctx context.Context) ([]models.Skills, error) {
	return services.SkillsRepository.FindAll(ctx)
}
