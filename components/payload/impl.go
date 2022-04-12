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
	paylaodResp         *payloadM.PayloadResp
	constraintComponent *constraint.Component
	feedbackComponent   *feedbackComp.Component
}

func New(
	paylaodResp *payloadM.PayloadResp,
	constraintComponent *constraint.Component,
	feedbackComponent *feedbackComp.Component,
) Component {

	return &impl{
		paylaodResp:         paylaodResp,
		constraintComponent: constraintComponent,
		feedbackComponent:   feedbackComponent,
	}
}

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*payloadM.PayloadEventResp, error) {

	payloadEventResp := &payloadM.PayloadEventResp{
		ID: im.paylaodResp.ID,
	}

	constraintEventResp, err := (*im.constraintComponent).Judge(ctx, e)
	if err != nil {
		return nil, err
	}

	payloadEventResp.ConstraintEventResp = constraintEventResp

	var feedReturn *feedbackM.FeedReturn

	if constraintEventResp.Pass {
		feedReturn, err = im.processFeedReturn(ctx, e, true)
		if err != nil {
			return nil, err
		}

		payloadEventResp.FeedReturn = feedReturn

		switch feedReturn.FeedReturnStatus {
		case feedbackM.ALL:
			payloadEventResp.PayloadEventJudgeType = payloadM.ALL
		case feedbackM.SOME:
			payloadEventResp.PayloadEventJudgeType = payloadM.SOME
		case feedbackM.NONE:
			payloadEventResp.PayloadEventJudgeType = payloadM.NONE
		}

	} else {
		feedReturn, err = im.processFeedReturn(ctx, e, false)
		if err != nil {
			return nil, err
		}

		payloadEventResp.FeedReturn = feedReturn

		payloadEventResp.PayloadEventJudgeType = payloadM.NONE
	}

	return payloadEventResp, nil
}

func (im *impl) processFeedReturn(ctx context.Context, e *eventM.Event, pass bool) (*feedbackM.FeedReturn, error) {

	// 計算回饋額
	feedReturn, err := (*im.feedbackComponent).Calculate(ctx, e, pass)
	if err != nil {
		return nil, err
	}

	return feedReturn, nil
}
