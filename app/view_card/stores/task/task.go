package task

import (
	"context"

	taskM "example.com/creditcard/app/view_card/models/task"
	"example.com/creditcard/app/view_card/utils/conn"
)

type Store interface {
	Create(ctx context.Context, conn *conn.Connection, task *taskM.Task) error
	CreateTasks(ctx context.Context, conn *conn.Connection, tasks []*taskM.Task) error
	UpdateByRewardID(ctx context.Context, conn *conn.Connection, tasks []*taskM.Task) error
	GetByRewardID(ctx context.Context, rewardID string) ([]*taskM.Task, error)
	DeleteByRewardID(ctx context.Context, conn *conn.Connection, rewardID string) error
}
