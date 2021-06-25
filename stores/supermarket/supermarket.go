package supermarket

import (
	"context"

	supermarketM "example.com/creditcard/models/supermarket"
)

type Store interface {
	Create(ctx context.Context, supermarket *supermarketM.Supermarket) error
	UpdateByID(ctx context.Context, supermarket *supermarketM.Supermarket) error
	GetAll(ctx context.Context) ([]*supermarketM.Supermarket, error)
}
