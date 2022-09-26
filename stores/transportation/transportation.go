package transportation

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, transportation *channel.Transportation) error
	UpdateByID(ctx context.Context, transportation *channel.Transportation) error
	GetAll(ctx context.Context) ([]*channel.Transportation, error)
	GetByID(ctx context.Context, ID string) (*channel.Transportation, error)
}
