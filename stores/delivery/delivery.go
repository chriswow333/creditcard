package delivery

import (
	"context"

	deliveryM "example.com/creditcard/models/delivery"
)

type Store interface {
	Create(ctx context.Context, delivery *deliveryM.Delivery) error
	UpdateByID(ctx context.Context, delivery *deliveryM.Delivery) error
	GetAll(ctx context.Context) ([]*deliveryM.Delivery, error)
	GetByID(ctx context.Context, ID string) (*deliveryM.Delivery, error)
}
