package reward

import (
	"context"

	rewardM "example.com/creditcard/models/reward"
)

type Store interface {
	Create(ctx context.Context, reward *rewardM.Reward) error
	GetByID(ctx context.Context, ID string) (*rewardM.Reward, error)
	GetByCardRewardID(ctx context.Context, cardID string) ([]*rewardM.Reward, error)
	UpdateByID(ctx context.Context, reward *rewardM.Reward) error
}
