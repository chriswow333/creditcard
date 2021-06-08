package card

import (
	"context"
	"time"

	cardM "example.com/creditcard/models/card"
	"example.com/creditcard/stores/card"
	"github.com/sirupsen/logrus"

	"go.uber.org/dig"
)

var (
	timeNow = time.Now
)

type impl struct {
	dig.In

	cardStore card.Store
}

func New(
	cardStore card.Store,
) Service {
	return &impl{
		cardStore: cardStore,
	}
}

func (im *impl) Create(ctx context.Context, card *cardM.Card) error {

	card.UpdateDate = timeNow().Unix()

	err := im.cardStore.Create(ctx, card)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (im *impl) GetByID(ctx context.Context, ID string) (*cardM.Card, error) {

	card, err := im.cardStore.GetByID(ctx, ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	return card, nil
}

func (im *impl) GetAll(ctx context.Context) ([]*cardM.Card, error) {

	cards, err := im.cardStore.GetAll(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}
	return cards, nil
}

func (im *impl) GetByBankID(ctx context.Context, bankID string) ([]*cardM.Card, error) {
	cards, err := im.cardStore.GetByBankID(ctx, bankID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}
	return cards, nil

}
