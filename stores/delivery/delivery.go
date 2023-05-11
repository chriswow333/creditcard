package delivery

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, delivery *channel.Delivery) error
	UpdateByID(ctx context.Context, delivery *channel.Delivery) error
	GetAll(ctx context.Context, offset, limit int) ([]*channel.Delivery, error)
	GetByID(ctx context.Context, ID string) (*channel.Delivery, error)
	FindLike(ctx context.Context, names []string) ([]*channel.Delivery, error)
}
