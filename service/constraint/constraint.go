package constraint

import (
	"context"

	constraintM "example.com/creditcard/models/constraint"
)

type Service interface {
	Create(ctx context.Context, constraint *constraintM.Constraint) error
	GetByID(ctxt context.Context, ID string) (*constraintM.Constraint, error)
	GetAll(ctx context.Context) ([]*constraintM.Constraint, error)
	GetByPrivilageID(ctx context.Context, privilageID string) ([]*constraintM.Constraint, error)
}
