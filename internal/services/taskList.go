package services

import (
	"context"
	"my_life/internal/domain"
	"my_life/internal/repository"
)

type TaskService struct {
	repo repository.Repository
}

func NewTaskService(repo repository.Repository) *TaskService {
	return &TaskService{repo: repo}
}

func (t TaskService) CreateList(ctx context.Context, list domain.TaskList) error {
	return t.repo.CreateList(ctx, &list)
}
