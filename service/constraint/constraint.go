package constraint

import (
	"context"

	constraintM "example.com/creditcard/models/constraint"
)

type Service interface {
	UpdateByRewardID(ctx context.Context, rewardID string, constraints []*constraintM.Constraint) error
	GetByRewardID(ctx context.Context, rewardID string) ([]*constraintM.Constraint, error)
}
