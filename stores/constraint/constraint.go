package constraint

import (
	"context"

	constraintM "example.com/creditcard/models/constraint"
)

type Store interface {
	Create(ctx context.Context, constraint *constraintM.Constraint) error
	GetByID(ctx context.Context, ID string) (*constraintM.Constraint, error)
	GetAll(ctx context.Context) ([]*constraintM.Constraint, error)
	GetByRewardID(ctx context.Context, rewardID string) ([]*constraintM.Constraint, error)
}
