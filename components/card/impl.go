package card

import (
	"context"
	"errors"

	"example.com/creditcard/components/reward"
	cardM "example.com/creditcard/models/card"
	eventM "example.com/creditcard/models/event"
	feedbackM "example.com/creditcard/models/feedback"
	rewardM "example.com/creditcard/models/reward"
	"github.com/sirupsen/logrus"
)

type impl struct {
	card         *cardM.Card
	rewardMapper map[rewardM.RewardType][]*reward.Component
}

func New(
	card *cardM.Card,
	rewardMapper map[rewardM.RewardType][]*reward.Component,
) Component {
	return &impl{
		card:         card,
		rewardMapper: rewardMapper,
	}
}

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*eventM.CardResp, error) {

	card := &eventM.CardResp{
		ID:         im.card.ID,
		BankID:     im.card.BankID,
		Name:       im.card.Name,
		StartDate:  im.card.StartDate,
		EndDate:    im.card.EndDate,
		UpdateDate: im.card.UpdateDate,
		LinkURL:    im.card.LinkURL,
	}

	cardRewards := []*eventM.CardReward{}

	if e.RewardType != 0 {
		// means using specficed type

	} else {
		for rewardType, rs := range im.rewardMapper {

			rewards := []*eventM.RewardResp{}

			for _, r := range rs {
				reward, err := (*r).Satisfy(ctx, e)

				if err != nil {
					logrus.WithFields(logrus.Fields{
						"": "",
					}).Error(err)
					return nil, err
				}
				rewards = append(rewards, reward)
			}

			cardReward, err := im.calculateCardReward(ctx, rewardType, rewards)
			if err != nil {
				return nil, err
			}
			cardRewards = append(cardRewards, cardReward)
		}
	}

	card.CardRewards = cardRewards

	return card, nil
}

func (im *impl) calculateCardReward(ctx context.Context, rewardType rewardM.RewardType, rewards []*eventM.RewardResp) (*eventM.CardReward, error) {

	var total float64

	cardReward := &eventM.CardReward{}

	for _, r := range rewards {

		if r.Constraint.Feedback == nil {
			return nil, errors.New("no feedback")
		}

		total = r.Constraint.Feedback.Total
		if r.Pass {
			switch rewardType {
			case rewardM.Cash:

				bonus, actualCashback := im.getCashBonus(ctx, r.Constraint.Feedback)

				cardReward.TotalGetBonus += bonus
				cardReward.TotalGetCash += actualCashback

			case rewardM.Point:

			default:

			}

		}
	}
	cardReward.TotalCost = total
	cardReward.Rewards = rewards
	return cardReward, nil

}

func (im *impl) getCashBonus(ctx context.Context, feedbackResp *feedbackM.Feedback) (float64, float64) {

	if feedbackResp.IsFeedbackGet {
		return feedbackResp.CashBack.CashBackLimit.Bonus, feedbackResp.CashBack.ActualCashBack
	} else {
		return 0.0, 0.0
	}

}
