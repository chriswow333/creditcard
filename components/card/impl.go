package card

import (
	"context"

	"example.com/creditcard/components/reward"
	eventM "example.com/creditcard/models/event"
	"github.com/sirupsen/logrus"
)

type impl struct {
	rewards []*reward.Component
}

func New(
	rewards []*reward.Component,
) Component {

	return &impl{
		rewards: rewards,
	}
}

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	resp := &eventM.Response{}

	for _, r := range im.rewards {
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
