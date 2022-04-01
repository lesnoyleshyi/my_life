package services

import "my_life/internal/repository"

type TaskService struct {
	repo repository.Repository
}

func NewTaskService(repo repository.Repository) *TaskService {
	return &TaskService{repo: repo}
}
