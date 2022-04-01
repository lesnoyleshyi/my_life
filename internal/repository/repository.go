package repository

import (
	"context"
	"github.com/jackc/pgx/v4"

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
	GetListsByUId(ctx context.Context, UId int) ([]domain.TaskList, error)
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserById(ctx context.Context, UId uint64) (*domain.User, error)
}

var TxOpts = pgx.TxOptions{
	IsoLevel:       pgx.ReadCommitted,
	AccessMode:     pgx.ReadOnly,
	DeferrableMode: pgx.Deferrable,
}
