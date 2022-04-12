package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
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

func (r sectionsRepo) GetSectionsByUId(ctx context.Context, UId int32) ([]domain.TaskSection, error) {
	return nil, nil
}
