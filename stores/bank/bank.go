package bank

import (
	"context"

	"example.com/creditcard/models/bank"
)

type Store interface {
	Create(ctx context.Context, bank *bank.Bank) error
	UpdateByID(ctx context.Context, bank *bank.Bank) error
	GetByID(ctx context.Context, ID string) (*bank.Bank, error)
	GetAll(ctx context.Context) ([]*bank.Bank, error)
}
