package hotel

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, hotel *channel.Hotel) error
	UpdateByID(ctx context.Context, hotel *channel.Hotel) error
	GetAll(ctx context.Context) ([]*channel.Hotel, error)
	GetByID(ctx context.Context, ID string) (*channel.Hotel, error)
	FindLike(ctx context.Context, names []string) ([]*channel.Hotel, error)
}
