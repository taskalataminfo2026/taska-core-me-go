package services

import (
	"context"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
	"taska-core-me-go/cmd/api/constants"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/repositories"
)

//go:generate mockgen -destination=../mocks/services/$GOFILE -package=mservices -source=./$GOFILE

type ITaskerServices interface {
	Find(ctx context.Context, request models.ParamsProfile) (models.Tasker, error)
}

type TaskerServices struct {
	SkillsRepository repositories.ISkillsRepository
}

func (services *TaskerServices) Find(ctx context.Context, request models.ParamsProfile) (models.Tasker, error) {
	logger.StandardInfo(ctx, constants.LayerService, constants.ModuleTasker, constants.FunctionTaskerFind, "Creando skill")
	return models.Tasker{}, nil
}
