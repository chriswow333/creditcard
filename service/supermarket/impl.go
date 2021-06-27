package supermarket

import (
	"context"

	supermarketM "example.com/creditcard/models/supermarket"
	supermarketStore "example.com/creditcard/stores/supermarket"
	"github.com/sirupsen/logrus"

	uuid "github.com/nu7hatch/gouuid"
)

type impl struct {
	supermarketStore supermarketStore.Store
}

func New(
	supermarketStore supermarketStore.Store,
) Service {
	return &impl{
		supermarketStore: supermarketStore,
	}
}

func (im *impl) Create(ctx context.Context, supermarket *supermarketM.Supermarket) error {

	ID, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}
	supermarket.ID = ID.String()

	if err := im.supermarketStore.Create(ctx, supermarket); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) UpdateByID(ctx context.Context, supermarket *supermarketM.Supermarket) error {
	if err := im.supermarketStore.UpdateByID(ctx, supermarket); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) GetAll(ctx context.Context) ([]*supermarketM.Supermarket, error) {
	supermarkets, err := im.supermarketStore.GetAll(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return supermarkets, nil
}
