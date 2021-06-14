package creditcardreward

import (
	"context"

	bankComp "example.com/creditcard/components/bank"
	cardComp "example.com/creditcard/components/card"
	constraintComp "example.com/creditcard/components/constraint"
	rewardComp "example.com/creditcard/components/reward"
	bankM "example.com/creditcard/models/bank"
	cardM "example.com/creditcard/models/card"
	constraintM "example.com/creditcard/models/constraint"
	rewardM "example.com/creditcard/models/reward"
)

type impl struct {
}

func New() Builder {
	return &impl{}

}

func (im *impl) NewCreditcard(ctx context.Context, settings []*bankM.Bank) ([]*bankComp.Component, error) {

	banks := []*bankComp.Component{}
	for _, b := range settings {
		cards := []*cardComp.Component{}

		for _, c := range b.Cards {

			rewards := []*rewardComp.Component{}

			for _, r := range c.Rewards {

				constraints := []*constraintComp.Component{}

				for _, co := range r.Constraints {
					constraint, _ := im.getConstraintComponent(ctx, co)
					constraints = append(constraints, constraint)
				}
				reward, _ := im.getRewardComponent(ctx, r, constraints)
				rewards = append(rewards, reward)
			}
			card, _ := im.getCardComponent(ctx, c, rewards)
			cards = append(cards, card)
		}
		bank, _ := im.getBankComponent(ctx, b, cards)
		banks = append(banks, bank)
	}

	return banks, nil
}

func (im *impl) getConstraintComponent(ctx context.Context, c *constraintM.Constraint) (*constraintComp.Component, error) {

	return nil, nil
}

func (im *impl) getRewardComponent(ctx context.Context, r *rewardM.Reward, constraints []*constraintComp.Component) (*rewardComp.Component, error) {

	return nil, nil
}

func (im *impl) getCardComponent(ctx context.Context, card *cardM.Card, rewards []*rewardComp.Component) (*cardComp.Component, error) {

	return nil, nil
}

func (im *impl) getBankComponent(ctx context.Context, bank *bankM.Bank, cards []*cardComp.Component) (*bankComp.Component, error) {

	return nil, nil
}
