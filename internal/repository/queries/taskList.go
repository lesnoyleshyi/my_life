package queries

import (
	"context"
	"fmt"
	"my_life/internal/domain"
)

const createTaskList = `INSERT INTO lists (UId, emoji, title, order_, relevanceTime)
						VALUES ($1, $2, $3, $4, $5);`

func (q *Queries) CreateList(ctx context.Context, l *domain.TaskList) error {

	tx, err := q.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	res, err := q.pool.Exec(ctx, createTaskList, l.UId, l.Emoji, l.Title, l.Order, l.RelevanceTime)
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

func (q *Queries) GetList(ctx context.Context, UId string) ([]domain.TaskList, error) {

	rows, err := q.pool.Query(ctx, selectByUId, UId)
	if err != nil {
		return nil, fmt.Errorf("error retrieving data from database, %w", err)
	}
	defer rows.Close()

	var lists []domain.TaskList

	for rows.Next() {
		var l domain.TaskList
		err := rows.Scan(&l.Emoji, &l.Title, &l.Order, &l.RelevanceTime)
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
