package feature

import (
	"context"

	cardM "example.com/creditcard/app/view_card/models/card"
)

type Store interface {
	CreateByCardID(ctx context.Context, cardID string, feature *cardM.Feature) error
	GetByCardID(ctx context.Context, cardID string) (*cardM.Feature, error)
	DeleteByCardID(ctx context.Context, cardID string) error
}
