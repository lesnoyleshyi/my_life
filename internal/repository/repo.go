package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"my_life/internal/domain"
	"my_life/internal/repository/queries"
)

type repo struct {
	*queries.Queries
	pool *pgxpool.Pool
}

func NewRepository(pgxPool *pgxpool.Pool) Repository {
	return &repo{
		Queries: queries.New(pgxPool),
		pool:    pgxPool,
	}
}

type Repository interface {
	CreateList(ctx context.Context, list *domain.TaskList) error
	GetListsById(ctx context.Context, UId int) ([]domain.TaskList, error)
}
