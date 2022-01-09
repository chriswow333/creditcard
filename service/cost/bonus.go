package cost

import (
	"context"

	costM "example.com/creditcard/models/cost"
)

type Service interface {
	UpdateByRewardID(ctx context.Context, rewardID string, cost *costM.Cost) error
	GetByRewardID(ctx context.Context, rewardID string) (*costM.Cost, error)
}
