package constraintpayload

import (
	"context"

	"example.com/creditcard/components/constraint"
	feedbackComp "example.com/creditcard/components/feedback"

	eventM "example.com/creditcard/models/event"
	feedbackM "example.com/creditcard/models/feedback"

	constraintM "example.com/creditcard/models/constraint"
)

type impl struct {
	constraints []*constraint.Component

	feedbackComponent *feedbackComp.Component

	operator constraintM.OperatorType
	name     string
	desc     string
}

func New(
	constraints []*constraint.Component,
	feedbackComponent *feedbackComp.Component,
	constraintPayload *constraintM.ConstraintPayload,
) constraint.Component {

	return &impl{
		constraints:       constraints,
		feedbackComponent: feedbackComponent,

		operator: constraintPayload.Operator,
		name:     constraintPayload.Name,
		desc:     constraintPayload.Desc,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.ConstraintResp, error) {

	constraint := &eventM.ConstraintResp{
		Name:           im.name,
		Desc:           im.desc,
		ConstraintType: constraintM.ConstraintPayloadType,
	}

	eventConstraints := []*eventM.ConstraintResp{}

	matches := 0
	misses := 0
	for _, c := range im.constraints {
		constraintResp, err := (*c).Judge(ctx, e)
		if err != nil {
			return nil, err
		}
		eventConstraints = append(eventConstraints, constraintResp)
		if constraintResp.Pass {
			matches++
		} else {
			misses++
		}
	}

	switch im.operator {
	case constraintM.OrOperator:
		if matches > 0 {
			constraint.Pass = true
		} else {
			constraint.Pass = false
		}
	case constraintM.AndOperator:
		if misses > 0 {
			constraint.Pass = false
		} else {
			constraint.Pass = true
		}
	}

	constraint.Constraints = eventConstraints

	if im.feedbackComponent != nil {

		var feedback *feedbackM.Feedback
		var err error

		if constraint.Pass {
			feedback, err = im.processFeedback(ctx, e, true)
			if err != nil {
				return nil, err
			}
		} else {
			feedback, err = im.processFeedback(ctx, e, false)

		}
		if err != nil {
			return nil, err
		}

		constraint.Feedback = feedback

		if !feedback.IsRewardGet {
			constraint.Pass = false
		}

	}

	return constraint, nil

}

func (im *impl) processFeedback(ctx context.Context, e *eventM.Event, pass bool) (*feedbackM.Feedback, error) {

	// 計算回饋額
	feedback, err := (*im.feedbackComponent).Calculate(ctx, e, pass)
	if err != nil {
		return nil, err
	}

	return feedback, nil
}
