package services

import (
	"context"
	"my_life/internal/domain"
	"my_life/internal/repository"
)

type ListService struct {
	repo repository.TaskList
}

func NewListService(repo repository.Repository) *ListService {
	return &ListService{repo: repo}
}

func (s ListService) CreateList(ctx context.Context, list *domain.TaskList) error {
	return s.repo.CreateList(ctx, list)
}

func (s ListService) GetListsByUId(ctx context.Context, UId int32) ([]domain.TaskList, error) {
	return s.repo.GetListsByUId(ctx, UId)
}
