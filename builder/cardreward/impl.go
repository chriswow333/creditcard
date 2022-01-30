package cardreward

import (
	"context"
	"errors"

	"go.uber.org/dig"

	cardComp "example.com/creditcard/components/card"
	constraintComp "example.com/creditcard/components/constraint"
	"example.com/creditcard/components/constraint/constraintpayload"
	"example.com/creditcard/components/constraint/customization"
	"example.com/creditcard/components/constraint/ecommerce"
	"example.com/creditcard/components/constraint/mobilepay"
	"example.com/creditcard/components/constraint/onlinegame"
	"example.com/creditcard/components/constraint/streaming"
	"example.com/creditcard/components/constraint/supermarket"
	"example.com/creditcard/components/constraint/timeinterval"
	feedbackComp "example.com/creditcard/components/feedback"
	cashBackComp "example.com/creditcard/components/feedback/cash_back"
	rewardComp "example.com/creditcard/components/reward"

	cardM "example.com/creditcard/models/card"
	constraintM "example.com/creditcard/models/constraint"
	feedbackM "example.com/creditcard/models/feedback"
	rewardM "example.com/creditcard/models/reward"
)

type impl struct {
	*dig.In
}

func New() Builder {

	return &impl{}

}

func (im *impl) BuildCardComponent(ctx context.Context, setting *cardM.Card) (*cardComp.Component, error) {

	rewardMapper := make(map[rewardM.RewardType][]*rewardComp.Component)
	for _, r := range setting.Rewards {

		constraint, err := im.getConstraintPayloadComponent(ctx, r.RewardType, r.ConstraintPayload)
		if err != nil {
			return nil, err
		}

		reward, _ := im.getRewardComponent(ctx, r, constraint)

		if _, ok := rewardMapper[r.RewardType]; ok {
			rewardMapper[r.RewardType] = append(rewardMapper[r.RewardType], reward)
		} else {
			rewardMapper[r.RewardType] = []*rewardComp.Component{reward}
		}
	}

	card, _ := im.getCardComponent(ctx, setting, rewardMapper)

	return card, nil
}

func (im *impl) getConstraintPayloadComponent(ctx context.Context, rewardType rewardM.RewardType, payload *constraintM.ConstraintPayload) (*constraintComp.Component, error) {

	var constraintComponents []*constraintComp.Component

	var constraintComponent constraintComp.Component
	if payload == nil {
		return nil, errors.New("no payload")
	}

	switch payload.ConstraintType {
	case constraintM.ConstraintPayloadType:

		for _, p := range payload.ConstraintPayloads {
			constraintComponentTemp, err := im.getConstraintPayloadComponent(ctx, rewardType, p)
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

	feedbackComponent, _ := im.getFeedbackComponent(ctx, rewardType, payload.Feedback)

	payloadCompoent := constraintpayload.New(constraintComponents, feedbackComponent, payload)

	return &payloadCompoent, nil
}

func (im *impl) getRewardComponent(ctx context.Context, r *rewardM.Reward, constraint *constraintComp.Component) (*rewardComp.Component, error) {
	component := rewardComp.New(r, constraint)
	return &component, nil
}

func (im *impl) getCardComponent(ctx context.Context, card *cardM.Card, rewardMapper map[rewardM.RewardType][]*rewardComp.Component) (*cardComp.Component, error) {
	component := cardComp.New(card, rewardMapper)
	return &component, nil
}

func (im *impl) getFeedbackComponent(ctx context.Context, rewardType rewardM.RewardType, feedback *feedbackM.Feedback) (*feedbackComp.Component, error) {

	if feedback == nil {
		return nil, nil
	}

	var feedbackComponent feedbackComp.Component

	switch rewardType {
	case rewardM.Cash:
		feedbackComponent = cashBackComp.New(
			feedback, feedback.CashBack,
		)
	case rewardM.Point:
		// costComponent = bonus.New()
	default:
		return nil, nil
	}

	return &feedbackComponent, nil
}
