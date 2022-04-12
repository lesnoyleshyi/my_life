package services

import (
	"context"
	"my_life/internal/domain"
	"my_life/internal/repository"
)

type sectionService struct {
	repo *repository.Repository
}

func newSectionService(repo *repository.Repository) *sectionService {
	return &sectionService{
		repo: repo,
	}
}

func (s sectionService) CreateSection(ctx context.Context, section *domain.TaskSection) error {
	return s.repo.CreateSection(ctx, section)
}

func (s sectionService) GetSectionsByUId(ctx context.Context, UId int32) ([]domain.TaskSection, error) {
	return s.repo.GetSectionsByUId(ctx, UId)
}
