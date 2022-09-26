package food

import (
	"context"

	"example.com/creditcard/models/channel"
)

type Store interface {
	Create(ctx context.Context, food *channel.Food) error
	UpdateByID(ctx context.Context, food *channel.Food) error
	GetAll(ctx context.Context) ([]*channel.Food, error)
	GetByID(ctx context.Context, ID string) (*channel.Food, error)
}
