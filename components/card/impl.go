package card

import (
	"context"

	"example.com/creditcard/components/reward"
	bonusM "example.com/creditcard/models/bonus"
	cardM "example.com/creditcard/models/card"
	eventM "example.com/creditcard/models/event"
	"github.com/sirupsen/logrus"
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

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*eventM.Card, error) {

	card := &eventM.Card{
		Name:      im.card.Name,
		Desc:      im.card.Desc,
		StartDate: im.card.StartDate,
		EndDate:   im.card.EndDate,
		LinkURL:   im.card.LinkURL,
	}
	rewards := []*eventM.Reward{}

	totalBonus := &bonusM.Bonus{
		BonusType: bonusM.Percentage,
	}

	countBonus := &bonusM.Bonus{
		BonusType: bonusM.Percentage,
	}

	for _, r := range im.rewards {
		reward, err := (*r).Satisfy(ctx, e)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}
		rewards = append(rewards, reward)
		if reward.Pass {
			countBonus.Point += reward.Bonus.Point
		}

		totalBonus.Point += reward.Bonus.Point
	}

	card.Rewards = rewards
	card.CountBonus = countBonus
	card.TotalBonus = totalBonus
	return card, nil
}
