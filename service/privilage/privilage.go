package privilage

import (
	"context"

	privilageM "example.com/creditcard/models/privilage"
)

type Service interface {
	Create(ctx context.Context, privilage *privilageM.Privilage) error
	GetByID(ctx context.Context, ID string) (*privilageM.Privilage, error)
	GetAll(ctx context.Context) ([]*privilageM.Privilage, error)
	GetByCardID(ctx context.Context, cardID string) ([]*privilageM.Privilage, error)
}
