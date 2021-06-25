package onlinegame

import (
	"context"

	onlinegameM "example.com/creditcard/models/onlinegame"
)

type Store interface {
	Create(ctx context.Context, onlinegame *onlinegameM.Onlinegame) error
	UpdateByID(ctx context.Context, onlinegame *onlinegameM.Onlinegame) error
	GetAll(ctx context.Context) ([]*onlinegameM.Onlinegame, error)
}
