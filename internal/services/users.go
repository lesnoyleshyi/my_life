package services

import (
	"context"
	"my_life/internal/domain"
	"my_life/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s UserService) GetUserById(ctx context.Context, UId int32) (*domain.User, error) {
	return s.repo.GetUserById(ctx, UId)
}
