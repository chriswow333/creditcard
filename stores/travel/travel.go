package travel

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, travel *channel.Travel) error
	UpdateByID(ctx context.Context, travel *channel.Travel) error
	GetAll(ctx context.Context, offset, limit int) ([]*channel.Travel, error)
	GetByID(ctx context.Context, ID string) (*channel.Travel, error)
	FindLike(ctx context.Context, names []string) ([]*channel.Travel, error)
}
