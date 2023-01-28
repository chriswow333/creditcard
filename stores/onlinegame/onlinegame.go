package onlinegame

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, onlinegame *channel.Onlinegame) error
	UpdateByID(ctx context.Context, onlinegame *channel.Onlinegame) error
	GetAll(ctx context.Context) ([]*channel.Onlinegame, error)
	GetByID(ctx context.Context, ID string) (*channel.Onlinegame, error)
	FindLike(ctx context.Context, names []string) ([]*channel.Onlinegame, error)
}
