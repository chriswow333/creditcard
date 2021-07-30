package reward

import (
	"context"

	rewardM "example.com/creditcard/app/view_card/models/reward"
	"example.com/creditcard/app/view_card/utils/conn"
)

type Store interface {
	Create(ctx context.Context, conn *conn.Connection, reward *rewardM.Reward) error
	GetByID(ctx context.Context, ID string) (*rewardM.Reward, error)
	UpdateByID(ctx context.Context, conn *conn.Connection, reward *rewardM.Reward) error
	GetByCardID(ctx context.Context, cardID string) ([]*rewardM.Reward, error)
}
