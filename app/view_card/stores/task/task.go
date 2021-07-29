package task

import (
	"context"

	taskM "example.com/creditcard/app/view_card/models/task"
)

type Store interface {
	Create(ctx context.Context, task *taskM.Task) error
	CreateTasks(ctx context.Context, tasks []*taskM.Task) error
	UpdateByRewardID(ctx context.Context, tasks []*taskM.Task) error
	GetByRewardID(ctx context.Context, rewardID string) ([]*taskM.Task, error)
}
