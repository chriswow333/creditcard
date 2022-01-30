package delivery

import (
	"context"

	deliveryM "example.com/creditcard/models/delivery"
	deliveryStore "example.com/creditcard/stores/delivery"
	"github.com/sirupsen/logrus"

	uuid "github.com/nu7hatch/gouuid"
)

type impl struct {
	deliveryStore deliveryStore.Store
}

func New(
	deliveryStore deliveryStore.Store,
) Service {
	return &impl{
		deliveryStore: deliveryStore,
	}
}

func (im *impl) Create(ctx context.Context, ecommerce *deliveryM.Delivery) error {
	ID, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}
	ecommerce.ID = ID.String()

	if err := im.deliveryStore.Create(ctx, ecommerce); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) UpdateByID(ctx context.Context, ecommerce *deliveryM.Delivery) error {
	if err := im.deliveryStore.UpdateByID(ctx, ecommerce); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) GetAll(ctx context.Context) ([]*deliveryM.Delivery, error) {
	ecommerces, err := im.deliveryStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return ecommerces, nil
}
