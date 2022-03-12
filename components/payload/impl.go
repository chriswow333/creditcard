package payload

import (
	"context"

	"example.com/creditcard/components/constraint"

	feedbackComp "example.com/creditcard/components/feedback"

	eventM "example.com/creditcard/models/event"
	feedbackM "example.com/creditcard/models/feedback"

	payloadM "example.com/creditcard/models/payload"
)

type impl struct {
	constraintComponent *constraint.Component
	feedbackComponent   *feedbackComp.Component
}

func New(
	constraintComponent *constraint.Component,
	feedbackComponent *feedbackComp.Component,
) Component {

	return &impl{
		constraintComponent: constraintComponent,
		feedbackComponent:   feedbackComponent,
	}
}

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*payloadM.PayloadResp, error) {

	payloadResp := &payloadM.PayloadResp{}

	constraintResp, err := (*im.constraintComponent).Judge(ctx, e)
	if err != nil {
		return nil, err
	}

	payloadResp.ConstraintResp = constraintResp
	payloadResp.Pass = constraintResp.Pass

	var feedReturn *feedbackM.FeedReturn

	if constraintResp.Pass {
		feedReturn, err = im.processFeedReturn(ctx, e, true)
		if err != nil {
			return nil, err
		}
	} else {
		feedReturn, err = im.processFeedReturn(ctx, e, false)
	}

	if err != nil {
		return nil, err
	}

	payloadResp.FeedReturn = feedReturn
	payloadResp.Feedback = (*im.feedbackComponent).GetFeedback(ctx)

	return payloadResp, nil
}

func (im *impl) processFeedReturn(ctx context.Context, e *eventM.Event, pass bool) (*feedbackM.FeedReturn, error) {

	// 計算回饋額
	feedReturn, err := (*im.feedbackComponent).Calculate(ctx, e, pass)
	if err != nil {
		return nil, err
	}

	return feedReturn, nil
}
