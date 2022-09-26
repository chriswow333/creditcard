package sport

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, sport *channel.Sport) error
	UpdateByID(ctx context.Context, sport *channel.Sport) error
	GetAll(ctx context.Context) ([]*channel.Sport, error)
	GetByID(ctx context.Context, ID string) (*channel.Sport, error)
}
