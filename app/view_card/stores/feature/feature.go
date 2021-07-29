package feature

import (
	"context"

	cardM "example.com/creditcard/app/view_card/models/card"
	"example.com/creditcard/app/view_card/utils/conn"
)

type Store interface {
	CreateByCardID(ctx context.Context, conn *conn.Connection, cardID string, feature *cardM.Feature) error
	GetByCardID(ctx context.Context, cardID string) (*cardM.Feature, error)
	DeleteByCardID(ctx context.Context, conn *conn.Connection, cardID string) error
}
