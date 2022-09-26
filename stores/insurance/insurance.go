package insurance

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, insurance *channel.Insurance) error
	UpdateByID(ctx context.Context, insurance *channel.Insurance) error
	GetAll(ctx context.Context) ([]*channel.Insurance, error)
	GetByID(ctx context.Context, ID string) (*channel.Insurance, error)
}
