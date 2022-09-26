package streaming

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, streaming *channel.Streaming) error
	UpdateByID(ctx context.Context, streaming *channel.Streaming) error
	GetAll(ctx context.Context) ([]*channel.Streaming, error)
	GetByID(ctx context.Context, ID string) (*channel.Streaming, error)
}
