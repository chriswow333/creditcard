package bank

import (
	"context"

	"example.com/creditcard/models/bank"
)

type Service interface {
	Create(ctx context.Context, card *bank.Bank) error
	UpdateByID(ctx context.Context, bank *bank.Bank) error
	GetByID(ctx context.Context, ID string) (*bank.Bank, error)
	GetRespByID(ctx context.Context, ID string) (*bank.BankResp, error)
	GetAll(ctx context.Context) ([]*bank.Bank, error)
	GetRespAll(ctx context.Context) ([]*bank.BankResp, error)
}
