package cardreward

import (
	"context"
	"fmt"

	cardComp "example.com/creditcard/components/card"
	constraintComp "example.com/creditcard/components/constraint"
	"example.com/creditcard/components/constraint/bonuslimit"
	"example.com/creditcard/components/constraint/constraintpayload"
	"example.com/creditcard/components/constraint/costlimit"
	"example.com/creditcard/components/constraint/customization"
	"example.com/creditcard/components/constraint/ecommerce"
	"example.com/creditcard/components/constraint/mobilepay"
	"example.com/creditcard/components/constraint/onlinegame"
	"example.com/creditcard/components/constraint/streaming"
	"example.com/creditcard/components/constraint/supermarket"
	"example.com/creditcard/components/constraint/timeinterval"
	rewardComp "example.com/creditcard/components/reward"
	cardM "example.com/creditcard/models/card"
	constraintM "example.com/creditcard/models/constraint"
	rewardM "example.com/creditcard/models/reward"
	"go.uber.org/dig"
)

type impl struct {
	*dig.In
}

func New() Builder {

	return &impl{}

}

func (im *impl) BuildCardComponent(ctx context.Context, setting *cardM.Card) (*cardComp.Component, error) {

	rewards := []*rewardComp.Component{}

	for _, r := range setting.Rewards {

		constraints := []*constraintComp.Component{}

		for _, co := range r.Constraints {

			constraint, _ := im.getConstraintComponent(ctx, co)
			constraints = append(constraints, constraint)
		}
		reward, _ := im.getRewardComponent(ctx, r, constraints)
		rewards = append(rewards, reward)
	}

	card, _ := im.getCardComponent(ctx, setting, rewards)
	return card, nil
}

func (im *impl) getConstraintComponent(ctx context.Context, c *constraintM.Constraint) (*constraintComp.Component, error) {

	constraintComponent, err := im.getConstraintPayloadComponent(ctx, c.ConstraintPayload)
	if err != nil {
		return nil, err
	}

	component := constraintComp.New(c, constraintComponent)

	return &component, nil
}

func (im *impl) getConstraintPayloadComponent(ctx context.Context, payload *constraintM.ConstraintPayload) (*constraintComp.Component, error) {

	var constraintComponents []*constraintComp.Component

	var constraintComponent constraintComp.Component
	fmt.Println(payload.ConstraintType)

	switch payload.ConstraintType {
	case constraintM.ConstraintPayloadType:

		for _, p := range payload.ConstraintPayloads {
			constraintComponentTemp, err := im.getConstraintPayloadComponent(ctx, p)
			if err != nil {
				return nil, err
			}
			constraintComponents = append(constraintComponents, constraintComponentTemp)
		}

	case constraintM.MobilepayType:
		constraintComponent = mobilepay.New(payload)
	case constraintM.EcommerceType:
		constraintComponent = ecommerce.New(payload)
	case constraintM.SupermarketType:
		constraintComponent = supermarket.New(payload)
	case constraintM.OnlinegameType:
		constraintComponent = onlinegame.New(payload)
	case constraintM.StreamingType:
		constraintComponent = streaming.New(payload)
	case constraintM.CostLimitType:
		constraintComponent = costlimit.New(payload)
	case constraintM.BonusLimitType:
		constraintComponent = bonuslimit.New(payload)
	case constraintM.TimeIntervalType:
		constraintComponent = timeinterval.New(payload)
	case constraintM.CustomizationType:
		constraintComponent = customization.New(payload)
	default:
		return nil, nil
	}

	if payload.ConstraintType != constraintM.ConstraintPayloadType {
		constraintComponents = append(constraintComponents, &constraintComponent)
	}

	payloadCompoent := constraintpayload.New(constraintComponents, payload)
	return &payloadCompoent, nil
}

func (im *impl) getRewardComponent(ctx context.Context, r *rewardM.Reward, constraints []*constraintComp.Component) (*rewardComp.Component, error) {

	component := rewardComp.New(r, constraints)
	return &component, nil
}

func (im *impl) getCardComponent(ctx context.Context, card *cardM.Card, rewards []*rewardComp.Component) (*cardComp.Component, error) {

	component := cardComp.New(card, rewards)
	return &component, nil
}
