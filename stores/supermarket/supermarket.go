package supermarket

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, supermarket *channel.Supermarket) error
	UpdateByID(ctx context.Context, supermarket *channel.Supermarket) error
	GetAll(ctx context.Context) ([]*channel.Supermarket, error)
	GetByID(ctx context.Context, ID string) (*channel.Supermarket, error)
	FindLike(ctx context.Context, names []string) ([]*channel.Supermarket, error)
}
