package publicutilities

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, publicUtility *channel.PublicUtility) error
	UpdateByID(ctx context.Context, publicUtility *channel.PublicUtility) error
	GetAll(ctx context.Context) ([]*channel.PublicUtility, error)
	GetByID(ctx context.Context, ID string) (*channel.PublicUtility, error)
	FindLike(ctx context.Context, names []string) ([]*channel.PublicUtility, error)
}
