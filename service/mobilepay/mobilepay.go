package mobilepay

import (
	"context"

	mobilepayM "example.com/creditcard/models/mobilepay"
)

type Service interface {
	Create(ctx context.Context, mobilepay *mobilepayM.Mobilepay) error
	UpdateByID(ctx context.Context, mobilepay *mobilepayM.Mobilepay) error
	GetAll(ctx context.Context) ([]*mobilepayM.Mobilepay, error)
}
