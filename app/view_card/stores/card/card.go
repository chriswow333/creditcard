package card

import (
	"context"

	cardM "example.com/creditcard/app/view_card/models/card"
	"example.com/creditcard/app/view_card/utils/conn"
)

type Store interface {
	Create(ctx context.Context, conn *conn.Connection, card *cardM.Card) error
	GetByID(ctx context.Context, ID string) (*cardM.Card, error)
	UpdateByID(ctx context.Context, conn *conn.Connection, card *cardM.Card) error
	GetAll(ctx context.Context) ([]*cardM.Card, error)
	GetByBankID(ctx context.Context, bankID string) ([]*cardM.Card, error)
}
