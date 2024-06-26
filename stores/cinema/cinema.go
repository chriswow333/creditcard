package cinema

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, cinema *channel.Cinema) error
	UpdateByID(ctx context.Context, cinema *channel.Cinema) error
	GetAll(ctx context.Context, offset, limit int) ([]*channel.Cinema, error)
	GetByID(ctx context.Context, ID string) (*channel.Cinema, error)
	FindLike(ctx context.Context, names []string) ([]*channel.Cinema, error)
}
