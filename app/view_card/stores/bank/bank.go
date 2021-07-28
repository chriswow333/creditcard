package bank

import (
	"context"

	bankM "example.com/creditcard/app/view_card/models/bank"
)

type Store interface {
	Create(ctx context.Context, bank *bankM.Bank) error
	GetByID(ctx context.Context, ID string) (*bankM.Bank, error)
	UpdateByID(ctx context.Context, bank *bankM.Bank) error
	GetAll(ctx context.Context) ([]*bankM.Bank, error)
}
