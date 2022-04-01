package queries

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"my_life/internal/domain"
)

var TxOpts = pgx.TxOptions{
	IsoLevel:       pgx.ReadCommitted,
	AccessMode:     pgx.ReadOnly,
	DeferrableMode: pgx.Deferrable,
}

const createUserQuery = `INSERT INTO users (id) VALUES ($1)`

func (q *Queries) CreateUser(ctx context.Context, user *domain.User) error {
	tx, err := q.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error creating transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	res, err := q.pool.Exec(ctx, createUserQuery, user.UId)
	if err != nil || res.RowsAffected() == 0 {
		return fmt.Errorf("error adding new user to database: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error commiting transaction: %w", err)
	}
	return nil
}

const getUserQuery = `SELECT id FROM users WHERE id=$1`

// GetUserById uses transaction only for educational purposes
func (q *Queries) GetUserById(ctx context.Context, UId uint64) (*domain.User, error) {
	tx, err := q.pool.BeginTx(ctx, TxOpts)
	if err != nil {
		return nil, fmt.Errorf("unable brgin transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	rows, err := q.pool.Query(ctx, getUserQuery, UId)
	if err != nil {
		return nil, fmt.Errorf("error retreiveing data from database: %w", err)
	}
	defer rows.Close()

	var user domain.User

	if err := rows.Scan(&user.UId); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("no user with id=%d is in database", UId)
		}
		return nil, fmt.Errorf("error scanning data to struct: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("error commiting transaction: %w", err)
	}

	return &user, nil
}
