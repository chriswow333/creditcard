package task

import (
	"context"

	"example.com/creditcard/models/task"
)

type Store interface {
	GetByCardID(ctx context.Context, rewardID string) ([]*task.Task, error)
	GetByID(ctx context.Context, ID string) (*task.Task, error)
	Create(ctx context.Context, task *task.Task) error
}
