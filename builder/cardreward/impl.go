package cardreward

import (
	"context"
	"errors"
	"fmt"

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

	cardM "example.com/creditcard/models/card"
	constraintM "example.com/creditcard/models/constraint"
	customizationM "example.com/creditcard/models/customization"
	"example.com/creditcard/models/event"
	feedbackM "example.com/creditcard/models/feedback"
	rewardM "example.com/creditcard/models/reward"

	customizationService "example.com/creditcard/service/customization"
)

type impl struct {
	*dig.In

	customizationService customizationService.Service
}

func New(
	customizationService customizationService.Service,
) Builder {
	return &impl{
		customizationService: customizationService,
	}

}

func (im *impl) BuildCardComponent(ctx context.Context, setting *cardM.Card) (*cardComp.Component, error) {

	rewardMapper := make(map[rewardM.RewardType][]*rewardComp.Component)
	payloadOperatorMapper := make(map[rewardM.RewardType]cardM.RewardOperator)

	for _, cr := range setting.CardRewards {

		rewardType := cr.RewardType
		payloadOperatorMapper[rewardType] = cr.RewardOperator

		for _, r := range cr.Rewards {
			payloadComponents := []*payloadComp.Component{}

			for _, p := range r.Payloads {
				constraintComponent, err := im.getConstraintComponent(ctx, p.Constraint)
				if err != nil {
					return nil, err
				}

				feedbackComponent, err := im.getFeedbackComponent(ctx, cr.RewardType, p.Feedback)
				if err != nil {
					return nil, err
				}

				payloadComponent := payloadComp.New(constraintComponent, feedbackComponent)
				payloadComponents = append(payloadComponents, &payloadComponent)
			}

			rewardComponent := rewardComp.New(cr.RewardType, r, payloadComponents)

			if _, ok := rewardMapper[cr.RewardType]; ok {
				rewardMapper[cr.RewardType] = append(rewardMapper[cr.RewardType], &rewardComponent)
			} else {
				rewardComponents := []*rewardComp.Component{}
				rewardComponents = append(rewardComponents, &rewardComponent)
				rewardMapper[cr.RewardType] = rewardComponents
			}

			fmt.Println("lennnnnnn")
			fmt.Println(cr.RewardType)
			fmt.Println(len(rewardMapper[cr.RewardType]))
		}
	}

	cardComponent := cardComp.New(setting, rewardMapper, payloadOperatorMapper)

	return &cardComponent, nil
}

func (im *impl) getConstraintComponent(ctx context.Context, constraint *constraintM.Constraint) (*constraintComp.Component, error) {

	constraintComponents := []*constraintComp.Component{}

	constraintType := constraint.ConstraintType

	var constraintComponent constraintComp.Component

	fmt.Println("in to get constraint component")
	switch constraintType {
	case constraintM.InnerConstraintType:
		fmt.Println("InnerConstraintType")
		for _, c := range constraint.InnerConstraints {
			constraintComponent, err := im.getConstraintComponent(ctx, c)
			if err != nil {
				return nil, err
			}
			fmt.Println("InnerConstraintType ", &constraintComponent)
			constraintComponents = append(constraintComponents, constraintComponent)
		}
		constraintComponent = constraintComp.New(constraintComponents, constraint)

	case constraintM.CustomizationType:
		fmt.Println("CustomizationType")
		customizations := []*customizationM.Customization{}

		for _, c := range constraint.Customizations {
			customization, err := im.customizationService.GetByID(ctx, c.ID)
			if err != nil {
				return nil, err
			}
			fmt.Println("customization ", &constraintComponent)
			customizations = append(customizations, customization)
		}

		constraint.Customizations = customizations

		constraintComponent = customizationComp.New(constraint)

	case constraintM.TimeIntervalType:
		fmt.Println("TimeIntervalType")
		constraintComponent = timeinterval.New(constraint)

	case constraintM.MobilepayType:
		fmt.Println("MobilepayType")
		constraintComponent = mobilepay.New(constraint)

	case constraintM.EcommerceType:
		fmt.Println("ecommerce")

		constraintComponent = ecommerce.New(constraint)
		fmt.Println(&constraintComponent)

	case constraintM.SupermarketType:
		constraintComponent = supermarket.New(constraint)

	case constraintM.OnlinegameType:
		fmt.Println("OnlinegameType")
		constraintComponent = onlinegame.New(constraint)
		fmt.Println(&constraintComponent)
		fmt.Println(constraintComponent.Judge(ctx, &event.Event{}))

	case constraintM.StreamingType:
		fmt.Println("StreamingType")
		constraintComponent = streaming.New(constraint)
		fmt.Println(&constraintComponent)
		fmt.Println(constraintComponent.Judge(ctx, &event.Event{}))

	default:
		return nil, errors.New("failed in mapping contraint type")

	}

	// constraintComponents = append(constraintComponents, &constraintComponent)

	// parentConstraintCompoent := constraintComp.New(constraintComponents, constraint)

	return &constraintComponent, nil
}

func (im *impl) getFeedbackComponent(ctx context.Context, rewardType rewardM.RewardType, feedback *feedbackM.Feedback) (*feedbackComp.Component, error) {

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
