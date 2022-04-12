package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"my_life/internal/domain"
)

type Authorisation interface {
	CreateUser(ctx context.Context, user *domain.User) (int, error)
	GetUser(ctx context.Context, username string, password string) (*domain.User, error)
}

type TaskList interface {
	CreateList(ctx context.Context, l *domain.TaskList) error
	GetListsByUId(ctx context.Context, UId int32) ([]domain.TaskList, error)
}

type TaskSection interface {
	CreateSection(ctx context.Context, s *domain.TaskSection) error
	GetSectionsByUId(ctx context.Context, UId int32) ([]domain.TaskSection, error)
}

type User interface {
	GetUserById(ctx context.Context, UId int32) (*domain.User, error)
}

type Repository struct {
	Authorisation
	TaskList
	TaskSection
	User
}

func NewRepository(pgxPool *pgxpool.Pool) *Repository {
	return &Repository{
		Authorisation: NewAuthPostgres(pgxPool),
		TaskList:      NewListRepo(pgxPool),
		TaskSection:   newSectionsRepo(pgxPool),
		User:          NewUserRepo(pgxPool),
	}
}
