package cardreward

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	cardComp "example.com/creditcard/components/card"
	constraintComp "example.com/creditcard/components/constraint"
	customizationComp "example.com/creditcard/components/constraint/customization"
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
	"example.com/creditcard/service/constraint"

	cardM "example.com/creditcard/models/card"
	constraintM "example.com/creditcard/models/constraint"
	customizationM "example.com/creditcard/models/customization"
	feedbackM "example.com/creditcard/models/feedback"
	rewardM "example.com/creditcard/models/reward"
)

type impl struct {
	*dig.In

	constraintService constraint.Service
}

func New(
	constraintService constraint.Service,
) Builder {
	return &impl{
		constraintService: constraintService,
	}

}

func (im *impl) BuildCardComponent(ctx context.Context, cardResp *cardM.CardResp) (cardComp.Component, error) {

	rewardMapper := make(map[rewardM.RewardType][]*rewardComp.Component)

	cardRewardOperatorMapper := make(map[rewardM.RewardType]cardM.CardRewardOperator)

	for _, cr := range cardResp.CardRewardResps {

		rewardType := cr.RewardType

		cardRewardOperatorMapper[rewardType] = cr.CardRewardOperator

		for _, r := range cr.RewardResps {

			payloadComponents := []*payloadComp.Component{}

			for _, p := range r.PayloadResps {
				constraintComponent, err := im.getConstraintComponent(ctx, p.ConstraintResp)
				if err != nil {
					logrus.New().Error(err)
					return nil, err
				}

				feedbackComponent, err := im.getFeedbackComponent(ctx, cr.RewardType, p.Feedback)
				if err != nil {
					logrus.New().Error(err)
					return nil, err
				}

				payloadComponent := payloadComp.New(p, constraintComponent, feedbackComponent)
				payloadComponents = append(payloadComponents, &payloadComponent)
			}

			rewardComponent := rewardComp.New(cr.RewardType, r, payloadComponents)

			if rewardCmp, ok := rewardMapper[cr.RewardType]; ok {
				rewardMapper[cr.RewardType] = append(rewardCmp, &rewardComponent)
			} else {
				rewardComponents := []*rewardComp.Component{}
				rewardComponents = append(rewardComponents, &rewardComponent)
				rewardMapper[cr.RewardType] = rewardComponents
			}

		}
	}

	cardComponent := cardComp.New(cardResp, rewardMapper, cardRewardOperatorMapper)

	return cardComponent, nil
}

func (im *impl) getConstraintComponent(ctx context.Context, constraintResp *constraintM.ConstraintResp) (*constraintComp.Component, error) {

	constraintType := constraintResp.ConstraintType

	var constraintComponent constraintComp.Component

	switch constraintType {
	case constraintM.InnerConstraintType:

		constraintComponents := []*constraintComp.Component{}

		for _, c := range constraintResp.InnerConstraints {
			constraintComponent, err := im.getConstraintComponent(ctx, c)
			if err != nil {
				return nil, err
			}
			constraintComponents = append(constraintComponents, constraintComponent)
		}

		constraintComponent = constraintComp.New(constraintComponents, constraintResp)

	case constraintM.CustomizationType:

		customizations := []*customizationM.Customization{}

		for _, c := range constraintResp.Customizations {
			customization, err := im.constraintService.GetCustomizationByID(ctx, c.ID)
			if err != nil {
				return nil, err
			}
			customizations = append(customizations, customization)
		}

		constraintResp.Customizations = customizations

		constraintComponent = customizationComp.New(constraintResp)

	case constraintM.TimeIntervalType:
		constraintComponent = timeinterval.New(constraintResp)

	case constraintM.MobilepayType:
		constraintComponent = mobilepay.New(constraintResp)

	case constraintM.EcommerceType:

		constraintComponent = ecommerce.New(constraintResp)

	case constraintM.SupermarketType:
		constraintComponent = supermarket.New(constraintResp)

	case constraintM.OnlinegameType:
		constraintComponent = onlinegame.New(constraintResp)

	case constraintM.StreamingType:
		constraintComponent = streaming.New(constraintResp)

	default:
		return nil, errors.New("failed in mapping contraint type")

	}

	return &constraintComponent, nil
}

func (im *impl) getFeedbackComponent(ctx context.Context, rewardType rewardM.RewardType, feedback *feedbackM.Feedback) (*feedbackComp.Component, error) {

	switch rewardType {
	case rewardM.CASH_TWD:
		cashbackComponent := cashbackComp.New(feedback.Cashback)
		return &cashbackComponent, nil
	case rewardM.POINT:
		return nil, nil
	default:
		return nil, nil
	}
}
