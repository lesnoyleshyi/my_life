package services

import (
	"context"
	"my_life/internal/repository"
	"net/http"
)

type Authorisation interface {
	repository.Authorisation
	GenerateToken(ctx context.Context, username, password string) (string, error)
	VerifyToken(next http.Handler) http.Handler
}

type TaskList interface {
	repository.TaskList
}

type User interface {
	repository.User
}

type TaskSection interface {
	repository.TaskSection
}

type Service struct {
	Authorisation
	TaskList
	TaskSection
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(*repo),
		TaskList:      NewListService(*repo),
		TaskSection:   newSectionService(repo),
		User:          NewUserService(*repo),
	}
}
