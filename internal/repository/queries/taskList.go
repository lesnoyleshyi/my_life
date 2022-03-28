package queries

import (
	"context"
	"my_life/internal/domain"
)

const createTaskList = `INSERT INTO TABLE lists () VALUES ()`

func (q *Queries) CreateList(ctx context.Context, list domain.TaskList) error {
	return nil
}
