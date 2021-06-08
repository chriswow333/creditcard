package bank

import (
	"context"
	"time"

	bankM "example.com/creditcard/models/bank"
	"example.com/creditcard/stores/bank"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

var (
	timeNow = time.Now
)

type impl struct {
	dig.In

	bankStore bank.Store
}

func New(
	bankStore bank.Store,
) Service {
	return &impl{
		bankStore: bankStore,
	}
}

func (im *impl) Create(ctx context.Context, bank *bankM.Bank) error {

	bank.UpdateDate = timeNow().Unix()

	err := im.bankStore.Create(ctx, bank)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (im *impl) GetByID(ctx context.Context, ID string) (*bankM.Bank, error) {

	bank, err := im.bankStore.GetByID(ctx, ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	return bank, nil
}

func (im *impl) GetAll(ctx context.Context) ([]*bankM.Bank, error) {

	banks, err := im.bankStore.GetAll(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}
	return banks, nil
}
