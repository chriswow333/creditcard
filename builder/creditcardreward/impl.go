package creditcardreward

import (
	"context"
	"fmt"

	bankComp "example.com/creditcard/components/bank"
	cardComp "example.com/creditcard/components/card"
	constraintComp "example.com/creditcard/components/constraint"
	"example.com/creditcard/components/constraint/accountbase"
	"example.com/creditcard/components/constraint/constraintpayload"
	"example.com/creditcard/components/constraint/ecommerce"
	"example.com/creditcard/components/constraint/mobilepay"
	"example.com/creditcard/components/constraint/moneybase"
	"example.com/creditcard/components/constraint/onlinegame"
	"example.com/creditcard/components/constraint/streaming"
	"example.com/creditcard/components/constraint/supermarket"
	timeBase "example.com/creditcard/components/constraint/timebase"
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

func (im *impl) getConstraintPayloadComponent(ctx context.Context, payload *constraintM.ConstraintPayload) (*constraintComp.Component, error) {

	var constraintComponents []*constraintComp.Component

	var constraintComponent constraintComp.Component
	switch payload.ConstraintType {
	case constraintM.ConstraintPayloadType:

		for _, p := range payload.ConstraintPayloads {
			constraintComponentTemp, _ := im.getConstraintPayloadComponent(ctx, p)
			constraintComponents = append(constraintComponents, constraintComponentTemp)
		}

	case constraintM.MobilepayType:
		constraintComponent = mobilepay.New(payload.Mobilepays, payload.Operator)
	case constraintM.EcommerceType:
		constraintComponent = ecommerce.New(payload.Ecommerces, payload.Operator)
	case constraintM.SupermarketType:
		constraintComponent = supermarket.New(payload.Supermarkets, payload.Operator)
	case constraintM.OnlinegameType:
		constraintComponent = onlinegame.New(payload.Onlinegames, payload.Operator)
	case constraintM.StreamingType:
		constraintComponent = streaming.New(payload.Streamings, payload.Operator)
	case constraintM.TimeBaseType:
		constraintComponent = timeBase.New(payload.TimeBases, payload.Operator)
	case constraintM.AccountBaseType:
		constraintComponent = accountbase.New(payload.AccountBases, payload.Operator)
	case constraintM.MoneyBaseType:
		constraintComponent = moneybase.New(payload.MoneyBases, payload.Operator)
	default:
		fmt.Println("error")
		break
	}

	if payload.ConstraintType != constraintM.ConstraintPayloadType {
		constraintComponents = append(constraintComponents, &constraintComponent)
	}

	payloadCompoent := constraintpayload.New(constraintComponents, payload.Operator)
	return &payloadCompoent, nil
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
