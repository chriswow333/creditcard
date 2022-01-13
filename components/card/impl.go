package card

import (
	"context"

	"github.com/sirupsen/logrus"

	"example.com/creditcard/components/reward"
	cardM "example.com/creditcard/models/card"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	card    *cardM.Card
	rewards []*reward.Component
}

func New(
	card *cardM.Card,
	rewards []*reward.Component,
) Component {

	return &impl{
		card:    card,
		rewards: rewards,
	}
}

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*eventM.CardResp, error) {

	card := &eventM.CardResp{
		Name:      im.card.Name,
		Desc:      im.card.Desc,
		StartDate: im.card.StartDate,
		EndDate:   im.card.EndDate,
		LinkURL:   im.card.LinkURL,
	}

	rewards := []*eventM.RewardResp{}

	for _, r := range im.rewards {

		reward, err := (*r).Satisfy(ctx, e)

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		rewards = append(rewards, reward)
	}

	card.Rewards = rewards
	return card, nil
}
