package bank

import (
	"context"

	"example.com/creditcard/components/card"
	eventM "example.com/creditcard/models/event"
	"github.com/sirupsen/logrus"
)

type impl struct {
	cards []*card.Component
}

func New(
	cards []*card.Component,
) Component {

	return &impl{
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
