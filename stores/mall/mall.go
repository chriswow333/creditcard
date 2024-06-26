package mall

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, mall *channel.Mall) error
	UpdateByID(ctx context.Context, mall *channel.Mall) error
	GetAll(ctx context.Context, offset, limit int) ([]*channel.Mall, error)
	GetByID(ctx context.Context, ID string) (*channel.Mall, error)
	FindLike(ctx context.Context, names []string) ([]*channel.Mall, error)
}
