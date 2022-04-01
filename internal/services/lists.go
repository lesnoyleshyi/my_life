package services

import (
	"context"
	"my_life/internal/domain"
)

func (t TaskService) CreateList(ctx context.Context, list *domain.TaskList) error {
	return t.repo.CreateList(ctx, list)
}

func (t TaskService) GetListsByUId(ctx context.Context, UId int) ([]domain.TaskList, error) {
	return t.repo.GetListsByUId(ctx, UId)
}
