package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"my_life/internal/domain"
)

type sectionsRepo struct {
	pool *pgxpool.Pool
}

func newSectionsRepo(pool *pgxpool.Pool) *sectionsRepo {
	return &sectionsRepo{
		pool: pool,
	}
}

const createSectionQuery string = `INSERT INTO sections (UId, listId, title, order_)
															VALUES ($1, $2, $3, $4);`

func (r sectionsRepo) CreateSection(ctx context.Context, s *domain.TaskSection) error {
	UId, ok := ctx.Value("UId").(int32)
	if !ok {
		return fmt.Errorf("can't retreive UId from context")
	}

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	res, err := tx.Exec(ctx, createSectionQuery, UId, s.ListId, s.Title, s.Order)
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

const getSectionsQuery = `SELECT id, UId, listId, title, order_, relevanceTime
							FROM sections WHERE UId = $1;`

func (r sectionsRepo) GetSectionsByUId(ctx context.Context, UId int32) ([]domain.TaskSection, error) {
	var sections []domain.TaskSection

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initialising transaction: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	rows, err := tx.Query(ctx, getSectionsQuery, UId)
	for rows.Next() {
		var r domain.TaskSection
		err := rows.Scan(&r.Id, &r.UId, &r.ListId, &r.Title, &r.Order, &r.RelevanceTime)
		if err != nil {
			log.Println(err)
		}
		sections = append(sections, r)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("error scanning rows from db to struct")
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("error commiting transaction: %w", err)
	}
	return sections, nil
}
