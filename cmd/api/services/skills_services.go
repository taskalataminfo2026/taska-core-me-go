package services

import (
	"context"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/repositories"
)

//go:generate mockgen -destination=../mocks/controllers/$GOFILE -package=mcontrollers -source=./$GOFILE

type ISkillsServices interface {
	SkillsSearch(ctx context.Context, request models.ParamsSkillsSearch) ([]models.SkillsResponse, error)
	SkillsList(ctx context.Context) ([]models.SkillsResponse, error)
}

type SkillsServices struct {
	SkillsRepository repositories.ISkillsRepository
}

func (services *SkillsServices) SkillsSearch(ctx context.Context, request models.ParamsSkillsSearch) ([]models.SkillsResponse, error) {
	return services.SkillsRepository.FindBy(ctx, request)
}

func (services *SkillsServices) SkillsList(ctx context.Context) ([]models.SkillsResponse, error) {
	return services.SkillsRepository.FindAll(ctx)
}
