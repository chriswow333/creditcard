package reward

import (
	"context"

	rewardM "example.com/creditcard/models/reward"
)

type Service interface {
	Create(ctx context.Context, reward *rewardM.Reward) error
	GetByID(ctx context.Context, ID string) (*rewardM.Reward, error)
	GetRespByID(ctx context.Context, ID string) (*rewardM.RewardResp, error)
	UpdateByID(ctx context.Context, reward *rewardM.Reward) error
	GetRespByCardID(ctx context.Context, cardID string) ([]*rewardM.RewardResp, error)
	GetByCardID(ctx context.Context, cardID string) ([]*rewardM.Reward, error)
}
