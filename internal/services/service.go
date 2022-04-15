package services

import (
	"context"
	"my_life/internal/domain"
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
	GetFullUserInfo(ctx context.Context, UId int32) (*domain.Reply, error)
}

type TaskSection interface {
	repository.TaskSection
}

type Tasker interface {
	repository.Tasker
}

type SubTasker interface {
	repository.SubTasker
}

type Service struct {
	Authorisation
	TaskList
	TaskSection
	Tasker
	SubTasker
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(*repo),
		TaskList:      NewListService(*repo),
		TaskSection:   newSectionService(repo),
		Tasker:        newTaskService(repo),
		SubTasker:     newSubtaskService(repo),
		User:          NewUserService(*repo),
	}
}
