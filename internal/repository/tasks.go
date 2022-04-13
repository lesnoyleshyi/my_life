package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"my_life/internal/domain"
)

type tasksRepo struct {
	pool *pgxpool.Pool
}

func newTaskRepo(pool *pgxpool.Pool) *tasksRepo {
	return &tasksRepo{
		pool: pool,
	}
}

const createTaskQuery = `INSERT INTO tasks (UId, listId, sectionId, title, isCompleted,
											completedDays, note, order_, repeatType,
											daysOfWeek, daysOfMonth, concreteDate,
											dateStart, dateEnd, dateReminder)
											VALUES ($1, $2, $3, $4, $5,
													$6, $7, $8, $9,
													$10, $11, $12,
													$13, $14, $15);`

func (r tasksRepo) CreateTask(ctx context.Context, t *domain.Task) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	res, err := tx.Exec(ctx, createTaskQuery, t.UId, t.ListId, t.SectionId, t.Title, t.IsCompleted,
		t.CompletedDays, t.Note, t.Order, t.RepeatType,
		t.DaysOfWeek, t.DaysOfMonth, t.ConcreteDate,
		t.DateStart, t.DateEnd, t.DateReminder)
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

const getListsByUId = `SELECT id, UId, listId, sectionId, title, isCompleted,
								completedDays, note, order_, repeatType,
								daysOfWeek, daysOfMonth, concreteDate,
								dateStart, dateEnd, dateReminder, relevanceTime
						FROM tasks WHERE UId = $1;`

func (r tasksRepo) GetTasksByUId(ctx context.Context, UId int32) ([]domain.Task, error) {
	var tasks []domain.Task

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	rows, err := tx.Query(ctx, getListsByUId, UId)
	for rows.Next() {
		var r domain.Task
		err := rows.Scan(&r.Id, &r.UId, &r.ListId, &r.SectionId, &r.Title, &r.IsCompleted,
			&r.CompletedDays, &r.Note, &r.Order, &r.RepeatType,
			&r.DaysOfWeek, &r.DaysOfMonth, &r.ConcreteDate,
			&r.DateStart, &r.DateEnd, &r.DateReminder, &r.RelevanceTime)
		if err != nil {
			log.Println(err)
		} else {
			tasks = append(tasks, r)
		}
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("error scanning rows from db to struct")
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("error commiting transaction: %w", err)
	}

	return tasks, nil
}
