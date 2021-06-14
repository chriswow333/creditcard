package constraint

import (
	"context"

	constraintM "example.com/creditcard/models/constraint"
)

type Service interface {
	Create(ctx context.Context, constraint *constraintM.Constraint) error
	GetByID(ctxt context.Context, ID string) (*constraintM.Constraint, error)
	GetAll(ctx context.Context) ([]*constraintM.Constraint, error)
	GetByRewardID(ctx context.Context, rewardID string) ([]*constraintM.Constraint, error)
}
