package amusement

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, amusement *channel.Amusement) error
	UpdateByID(ctx context.Context, amusement *channel.Amusement) error
	GetAll(ctx context.Context, offset, limit int) ([]*channel.Amusement, error)
	GetByID(ctx context.Context, ID string) (*channel.Amusement, error)
	FindLike(ctx context.Context, names []string) ([]*channel.Amusement, error)
}
