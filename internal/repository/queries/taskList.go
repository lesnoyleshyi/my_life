package queries

import (
	"context"
	"fmt"
	"my_life/internal/domain"
)

const createTaskList = `INSERT INTO TABLE lists (emoji, title, order_, relevance_time)
						VALUES ($1, $2, $3, $4)`

func (q *Queries) CreateList(ctx context.Context, l *domain.TaskList) error {

	tx, err := q.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	res, err := q.pool.Exec(ctx, createTaskList, l.Emoji, l.Title, l.Order, l.RelevanceTime)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("no rows were affected")
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}
