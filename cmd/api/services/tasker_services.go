package services

import (
	"context"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/repositories"
)

//go:generate mockgen -destination=../mocks/controllers/$GOFILE -package=mcontrollers -source=./$GOFILE

type ITaskerServices interface {
	GetTasker(ctx context.Context, request models.ParamsProfile) (models.Tasker, error)
}

type TaskerServices struct {
	SkillsRepository repositories.ISkillsRepository
}

func (services *TaskerServices) GetTasker(ctx context.Context, request models.ParamsProfile) (models.Tasker, error) {
	return models.Tasker{}, nil
}
