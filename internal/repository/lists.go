package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"my_life/internal/domain"
)

type listRepo struct {
	pool *pgxpool.Pool
}

func NewListRepo(pool *pgxpool.Pool) *listRepo {
	return &listRepo{pool: pool}
}

const createTaskList = `INSERT INTO lists (UId, emoji, title, order_)
						VALUES ($1, $2, $3, $4);`

func (r *listRepo) CreateList(ctx context.Context, tl *domain.TaskList) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	res, err := r.pool.Exec(ctx, createTaskList, tl.UId, tl.Emoji, tl.Title, tl.Order)
	if err != nil {
		return fmt.Errorf("error adding data to database: %w", err)
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("no rows were affected")
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("error commiting transaction: %w", err)
	}
	return nil
}

const selectByUId = `SELECT id, UId, emoji, title, order_, relevanceTime FROM lists WHERE UId=$1;`

func (r *listRepo) GetListsByUId(ctx context.Context, UId int32) ([]domain.TaskList, error) {
	rows, err := r.pool.Query(ctx, selectByUId, UId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving data from database, %w", err)
	}
	defer rows.Close()

	var lists []domain.TaskList

	for rows.Next() {
		var l domain.TaskList
		err := rows.Scan(&l.Id, &l.UId, &l.Emoji, &l.Title, &l.Order, &l.RelevanceTime)
		if err != nil {
			return nil, fmt.Errorf("error scanning data from row to struct: %w", err)
		}
		lists = append(lists, l)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error scanning data from row to struct: %w", err)
	}

	return lists, nil
}
