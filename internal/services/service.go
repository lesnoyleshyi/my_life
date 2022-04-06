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

type Service struct {
	Authorisation
	TaskList
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(*repo),
		TaskList:      NewListService(*repo),
		User:          NewUserService(*repo),
	}
}
