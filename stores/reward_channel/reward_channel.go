package reward_channel

import (
	"context"

	"example.com/creditcard/models/reward_channel"
)

type Store interface {
	Create(ctx context.Context, rewardConstraint *reward_channel.RewardChannel) error
	GetByRewardID(ctx context.Context, cardRewardID string) ([]*reward_channel.RewardChannel, error)
}
