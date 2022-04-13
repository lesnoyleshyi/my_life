package services

import (
	"context"
	"my_life/internal/domain"
	"my_life/internal/repository"
)

type subtaskService struct {
	repo *repository.Repository
}

func newSubtaskService(repo *repository.Repository) *subtaskService {
	return &subtaskService{
		repo: repo,
	}
}

func (s subtaskService) CreateSubTask(ctx context.Context, t *domain.Subtask) error {
	return s.repo.CreateSubTask(ctx, t)
}

func (s subtaskService) GetSubTasksByUId(ctx context.Context, UId int32) ([]domain.Subtask, error) {
	return s.repo.GetSubTasksByUId(ctx, UId)
}
