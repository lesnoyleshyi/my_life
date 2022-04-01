package services

import (
	"context"
	"my_life/internal/domain"
)

func (t TaskService) CreateUser(ctx context.Context, user *domain.User) error {
	return t.repo.CreateUser(ctx, user)
}

func (t TaskService) GetUserById(ctx context.Context, UId uint64) (*domain.User, error) {
	return t.repo.GetUserById(ctx, UId)
}
