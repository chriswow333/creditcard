package reward

import (
	"context"

	rewardM "example.com/creditcard/app/view_card/models/reward"
)

type Store interface {
	Create(ctx context.Context, reward *rewardM.Reward) error
	GetByID(ctx context.Context, ID string) (*rewardM.Reward, error)
	UpdateByID(ctx context.Context, reward *rewardM.Reward) error
	GetByCardID(ctx context.Context, cardID string) ([]*rewardM.Reward, error)
}
