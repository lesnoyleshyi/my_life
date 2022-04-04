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

const createTaskList = `INSERT INTO lists (UId, emoji, title, order_, relevanceTime)
						VALUES ($1, $2, $3, $4, $5);`

func (l *listRepo) CreateList(ctx context.Context, lll *domain.TaskList) error {
	tx, err := l.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	res, err := l.pool.Exec(ctx, createTaskList, lll.UId, lll.Emoji, lll.Title, lll.Order, lll.RelevanceTime)
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

const selectByUId = `SELECT * FROM lists WHERE UId=$1;`

func (l *listRepo) GetListsByUId(ctx context.Context, UId int64) ([]domain.TaskList, error) {
	rows, err := l.pool.Query(ctx, selectByUId, UId)
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
