package card

import (
	"context"

	cardM "example.com/creditcard/app/view_card/models/card"
)

type Service interface {
	Create(ctx context.Context, cardRepr *cardM.Repr) error
	GetByID(ctx context.Context, ID string) (*cardM.Repr, error)
	UpdateByID(ctx context.Context, cardRepr *cardM.Repr) error
	GetAll(ctx context.Context) ([]*cardM.Repr, error)
	GetByBankID(ctx context.Context, bankID string) ([]*cardM.Repr, error)
}
