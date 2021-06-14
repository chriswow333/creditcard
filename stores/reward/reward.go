package reward

import (
	"context"

	rewardM "example.com/creditcard/models/reward"
)

type Store interface {
	Create(ctx context.Context, reward *rewardM.Reward) error
	GetByID(ctx context.Context, ID string) (*rewardM.Reward, error)
	GetAll(ctx context.Context) ([]*rewardM.Reward, error)
	GetByCardID(ctx context.Context, cardID string) ([]*rewardM.Reward, error)
}
