package mobilepay

import (
	"context"

	mobilepayM "example.com/creditcard/models/mobilepay"
	mobilepayStore "example.com/creditcard/stores/mobilepay"
	"github.com/sirupsen/logrus"
)

type impl struct {
	mobilepayStore mobilepayStore.Store
}

func New(
	mobilepayStore mobilepayStore.Store,
) Service {
	return &impl{
		mobilepayStore: mobilepayStore,
	}
}

func (im *impl) Create(ctx context.Context, mobilepay *mobilepayM.Mobilepay) error {
	if err := im.mobilepayStore.Create(ctx, mobilepay); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) UpdateByID(ctx context.Context, mobilepay *mobilepayM.Mobilepay) error {
	if err := im.mobilepayStore.UpdateByID(ctx, mobilepay); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) GetAll(ctx context.Context) ([]*mobilepayM.Mobilepay, error) {
	mobilepays, err := im.mobilepayStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return mobilepays, nil
}
