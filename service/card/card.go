package card

import (
	"context"

	cardM "example.com/creditcard/models/card"
)

type Service interface {
	Create(ctx context.Context, card *cardM.Card) error
	GetByID(ctx context.Context, ID string) (*cardM.Card, error)
	GetAll(ctx context.Context) ([]*cardM.Card, error)
	GetByBankID(ctx context.Context, bankID string) ([]*cardM.Card, error)
}
