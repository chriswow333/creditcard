package cardreward

import (
	"context"

	"go.uber.org/dig"

	cardComp "example.com/creditcard/components/card"
	constraintComp "example.com/creditcard/components/constraint"
	"example.com/creditcard/components/constraint/ecommerce"
	"example.com/creditcard/components/constraint/mobilepay"
	"example.com/creditcard/components/constraint/onlinegame"
	"example.com/creditcard/components/constraint/streaming"
	"example.com/creditcard/components/constraint/supermarket"
	"example.com/creditcard/components/constraint/timeinterval"
	feedbackComp "example.com/creditcard/components/feedback"
	cashbackComp "example.com/creditcard/components/feedback/cashback"
	payloadComp "example.com/creditcard/components/payload"
	rewardComp "example.com/creditcard/components/reward"

	cardM "example.com/creditcard/models/card"
	constraintM "example.com/creditcard/models/constraint"
	"example.com/creditcard/models/feedback"
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
		payloadComponents := []*payloadComp.Component{}
		for _, p := range r.Payloads {
			constraintComponent, err := im.getConstraintComponent(ctx, p.Constraint)
			if err != nil {
				return nil, err
			}

			feedbackComponent, err := im.getFeedbackComponent(ctx, r.RewardType, p.Feedback)
			if err != nil {
				return nil, err
			}

			payloadComponent := payloadComp.New(constraintComponent, feedbackComponent)
			payloadComponents = append(payloadComponents, &payloadComponent)
		}

		rewardComponent := rewardComp.New(r, payloadComponents)

		if rewardComponents, ok := rewardMapper[r.RewardType]; ok {
			rewardComponents = append(rewardComponents, &rewardComponent)
		} else {
			rewardComponents := []*rewardComp.Component{}
			rewardComponents = append(rewardComponents, &rewardComponent)
			rewardMapper[r.RewardType] = rewardComponents
		}
	}

	cardComponent := cardComp.New(setting, rewardMapper, nil)

	return &cardComponent, nil
}

func (im *impl) getConstraintComponent(ctx context.Context, constraint *constraintM.Constraint) (*constraintComp.Component, error) {

	constraintComponents := []*constraintComp.Component{}

	constraintType := constraint.ConstraintType

	var constraintComponent constraintComp.Component

	switch constraintType {
	case constraintM.InnerConstraintType:
		for _, c := range constraint.InnerConstraints {
			constraintComponent, err := im.getConstraintComponent(ctx, c)
			if err != nil {
				return nil, err
			}
			constraintComponents = append(constraintComponents, constraintComponent)
		}

	case constraintM.CustomizationType:
		constraintComponent = timeinterval.New(constraint)

	case constraintM.TimeIntervalType:
		constraintComponent = timeinterval.New(constraint)

	case constraintM.MobilepayType:
		constraintComponent = mobilepay.New(constraint)

	case constraintM.EcommerceType:
		constraintComponent = ecommerce.New(constraint)

	case constraintM.SupermarketType:
		constraintComponent = supermarket.New(constraint)

	case constraintM.OnlinegameType:
		constraintComponent = onlinegame.New(constraint)

	case constraintM.StreamingType:
		constraintComponent = streaming.New(constraint)

	}

	constraintComponents = append(constraintComponents, &constraintComponent)

	parentConstraintCompoent := constraintComp.New(constraintComponents, constraint)

	return &parentConstraintCompoent, nil
}

func (im *impl) getFeedbackComponent(ctx context.Context, rewardType rewardM.RewardType, feedback *feedback.Feedback) (*feedbackComp.Component, error) {

	switch rewardType {
	case rewardM.InCash:
		cashbackComponent := cashbackComp.New(feedback.Cashback)
		return &cashbackComponent, nil
	case rewardM.OutCash:
		cashbackComponent := cashbackComp.New(feedback.Cashback)
		return &cashbackComponent, nil
	case rewardM.Point:
		return nil, nil
	default:
		return nil, nil
	}
}
