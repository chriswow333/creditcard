package ecommerce

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, ecommerce *channel.Ecommerce) error
	UpdateByID(ctx context.Context, ecommerce *channel.Ecommerce) error
	GetAll(ctx context.Context, offset, limit int) ([]*channel.Ecommerce, error)
	GetByID(ctx context.Context, ID string) (*channel.Ecommerce, error)
	FindLike(ctx context.Context, names []string) ([]*channel.Ecommerce, error)
}
