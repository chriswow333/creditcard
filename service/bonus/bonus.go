package bonus

import (
	"context"

	bonusM "example.com/creditcard/models/bonus"
)

type Service interface {
	UpdateByRewardID(ctx context.Context, rewardID string, bonus *bonusM.Bonus) error
	GetByRewardID(ctx context.Context, rewardID string) (*bonusM.Bonus, error)
}
