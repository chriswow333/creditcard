package appstore

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, appstore *channel.AppStore) error
	UpdateByID(ctx context.Context, appstore *channel.AppStore) error
	GetAll(ctx context.Context) ([]*channel.AppStore, error)
	GetByID(ctx context.Context, ID string) (*channel.AppStore, error)
	FindLike(ctx context.Context, names []string) ([]*channel.AppStore, error)
}
