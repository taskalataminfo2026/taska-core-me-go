package services

import (
	"context"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/repositories"
)

//go:generate mockgen -destination=../mocks/services/$GOFILE -package=mservices -source=./$GOFILE

type ICategoriesServices interface {
	List(ctx context.Context) ([]models.Category, error)
}

type CategoriesServices struct {
	CategoriesRepository repositories.ICategoriesRepository
}

func (services *CategoriesServices) List(ctx context.Context) ([]models.Category, error) {
	return services.CategoriesRepository.FindAll(ctx)
}
