package ecommerce

import (
	"context"

	"example.com/creditcard/models/ecommerce"
	ecommerceM "example.com/creditcard/models/ecommerce"
	ecommerceStore "example.com/creditcard/stores/ecommerce"
	"github.com/sirupsen/logrus"
)

type impl struct {
	ecommerceStore ecommerceStore.Store
}

func New(
	ecommerceStore ecommerceStore.Store,
) Service {
	return &impl{
		ecommerceStore: ecommerceStore,
	}
}

func (im *impl) Create(ctx context.Context, ecommerce *ecommerceM.Ecommerce) error {
	if err := im.ecommerceStore.Create(ctx, ecommerce); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) UpdateByID(ctx context.Context, ecommerce *ecommerce.Ecommerce) error {
	if err := im.ecommerceStore.UpdateByID(ctx, ecommerce); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) GetAll(ctx context.Context) ([]*ecommerceM.Ecommerce, error) {
	ecommerces, err := im.ecommerceStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return ecommerces, nil
}
