package services

import (
	"my_life/internal/repository"
)

type Authorisation interface {
	repository.Authorisation
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
