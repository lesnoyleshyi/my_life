package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"my_life/internal/domain"
)

type subtaskRepo struct {
	pool *pgxpool.Pool
}

func newSubtaskRepo(pool *pgxpool.Pool) *subtaskRepo {
	return &subtaskRepo{
		pool: pool,
	}
}

const createSubTaskQuery = `INSERT INTO subtasks (UId, listId, sectionId, taskId,
													title, isCompleted, order_)
											VALUES ($1, $2, $3, $4,
													$5, $6, $7);`

func (r subtaskRepo) CreateSubTask(ctx context.Context, st *domain.Subtask) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	res, err := tx.Exec(ctx, createSubTaskQuery, st.UId, st.ListId, st.SectionId, st.TaskId,
		st.Title, st.IsCompleted, st.Order)
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

const getSubTasksQuery = `SELECT id, UId, listId, sectionId, taskId, title,
										isCompleted, order_, relevanceTime
											FROM subtasks WHERE UId = $1;`

func (r subtaskRepo) GetSubTasksByUId(ctx context.Context, UId int32) ([]domain.Subtask, error) {
	var subtasks []domain.Subtask

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	rows, err := tx.Query(ctx, getSubTasksQuery, UId)
	for rows.Next() {
		var r domain.Subtask
		err := rows.Scan(&r.Id, &r.UId, &r.ListId, &r.SectionId, &r.TaskId, &r.Title,
			&r.IsCompleted, &r.Order, &r.RelevanceTime)
		if err != nil {
			log.Println(err)
		} else {
			subtasks = append(subtasks, r)
		}
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("error scanning rows from db to struct")
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("error commiting transaction: %w", err)
	}

	return subtasks, nil
}
