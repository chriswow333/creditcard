package bank

import (
	"context"

	bankM "example.com/creditcard/models/bank"
)

type Service interface {
	Create(ctx context.Context, card *bankM.Bank) error
	GetByID(ctx context.Context, ID string) (*bankM.Bank, error)
	GetAll(ctx context.Context) ([]*bankM.Bank, error)
}
