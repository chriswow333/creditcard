package transportation

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, transportation *channel.Transportation) error
	UpdateByID(ctx context.Context, transportation *channel.Transportation) error
	GetAll(ctx context.Context, offset, limit int) ([]*channel.Transportation, error)
	GetByID(ctx context.Context, ID string) (*channel.Transportation, error)
	FindLike(ctx context.Context, names []string) ([]*channel.Transportation, error)
}
