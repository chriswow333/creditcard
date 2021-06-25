package ecommerce

import (
	"context"

	"example.com/creditcard/models/ecommerce"
	ecommerceM "example.com/creditcard/models/ecommerce"
)

type Service interface {
	Create(ctx context.Context, ecommerce *ecommerceM.Ecommerce) error
	UpdateByID(ctx context.Context, ecommerce *ecommerce.Ecommerce) error
	GetAll(ctx context.Context) ([]*ecommerceM.Ecommerce, error)
}
