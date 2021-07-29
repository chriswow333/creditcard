package bank

import (
	"context"

	bankM "example.com/creditcard/app/view_card/models/bank"
)

type Service interface {
	Create(ctx context.Context, bank *bankM.Repr) error
	GetByID(ctx context.Context, ID string) (*bankM.Repr, error)
	UpdateByID(ctx context.Context, bank *bankM.Repr) error
	GetAll(ctx context.Context) ([]*bankM.Repr, error)
}
