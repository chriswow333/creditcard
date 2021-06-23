package bank

import (
	"context"

	"github.com/sirupsen/logrus"

	"example.com/creditcard/components/card"
	bankM "example.com/creditcard/models/bank"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	bank  *bankM.Bank
	cards []*card.Component
}

func New(
	bank *bankM.Bank,
	cards []*card.Component,
) Component {

	return &impl{
		bank:  bank,
		cards: cards,
	}
}

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	resp := &eventM.Response{}

	for _, r := range im.cards {
		_, err := (*r).Satisfy(ctx, e)

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}
	}

	return resp, nil
}
