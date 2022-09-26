package conveniencestore

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, convenienceStore *channel.ConvenienceStore) error
	UpdateByID(ctx context.Context, convenienceStore *channel.ConvenienceStore) error
	GetAll(ctx context.Context) ([]*channel.ConvenienceStore, error)
	GetByID(ctx context.Context, ID string) (*channel.ConvenienceStore, error)
}
