package delivery

import (
	"context"

	deliveryM "example.com/creditcard/models/delivery"
)

type Service interface {
	Create(ctx context.Context, delivery *deliveryM.Delivery) error
	UpdateByID(ctx context.Context, deliveryM *deliveryM.Delivery) error
	GetAll(ctx context.Context) ([]*deliveryM.Delivery, error)
}
