package insurance

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, insurance *channel.Insurance) error
	UpdateByID(ctx context.Context, insurance *channel.Insurance) error
	GetAll(ctx context.Context, offset, limit int) ([]*channel.Insurance, error)
	GetByID(ctx context.Context, ID string) (*channel.Insurance, error)
	FindLike(ctx context.Context, names []string) ([]*channel.Insurance, error)
}
