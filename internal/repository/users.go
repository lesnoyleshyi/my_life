package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"my_life/internal/domain"
)

type userRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *userRepo {
	return &userRepo{pool: pool}
}

var TxOpts = pgx.TxOptions{
	IsoLevel:       pgx.ReadCommitted,
	AccessMode:     pgx.ReadOnly,
	DeferrableMode: pgx.Deferrable,
}

const getUserIdQuery = `SELECT id FROM users WHERE id=$1`

// GetUserById uses transaction only for educational purposes
func (u *userRepo) GetUserById(ctx context.Context, UId int32) (*domain.User, error) {
	tx, err := u.pool.BeginTx(ctx, TxOpts)
	if err != nil {
		return nil, fmt.Errorf("unable brgin transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	row := u.pool.QueryRow(ctx, getUserIdQuery, UId)
	var user domain.User

	if err := row.Scan(&user.UId); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("there is no user with id=%d in database", UId)
		}
		return nil, fmt.Errorf("error scanning data to struct: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("error commiting transaction: %w", err)
	}

	return &user, nil
}
