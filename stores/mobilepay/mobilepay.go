package mobilepay

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, mobilepay *channel.Mobilepay) error
	UpdateByID(ctx context.Context, mobilepay *channel.Mobilepay) error
	GetAll(ctx context.Context, offset, limit int) ([]*channel.Mobilepay, error)
	GetByID(ctx context.Context, ID string) (*channel.Mobilepay, error)
	FindLike(ctx context.Context, names []string) ([]*channel.Mobilepay, error)
}
