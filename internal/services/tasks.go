package services

import (
	"context"
	"my_life/internal/domain"
	"my_life/internal/repository"
)

type taskService struct {
	repo *repository.Repository
}

func newTaskService(repo *repository.Repository) *taskService {
	return &taskService{
		repo: repo,
	}
}

func (s taskService) CreateTask(ctx context.Context, t *domain.Task) error {
	return s.repo.CreateTask(ctx, t)
}

func (s taskService) GetTasksByUId(ctx context.Context, UId int32) ([]domain.Task, error) {
	return s.repo.GetTasksByUId(ctx, UId)
}
