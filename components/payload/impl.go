package payload

import (
	"context"

	"example.com/creditcard/components/channel"
	"github.com/sirupsen/logrus"

	feedbackComp "example.com/creditcard/components/feedback"

	eventM "example.com/creditcard/models/event"
	feedbackM "example.com/creditcard/models/feedback"

	payloadM "example.com/creditcard/models/payload"
)

type impl struct {
	payload           *payloadM.Payload
	channelComponent  *channel.Component
	feedbackComponent *feedbackComp.Component
}

func New(
	payload *payloadM.Payload,
	channelComponent *channel.Component,
	feedbackComponent *feedbackComp.Component,
) Component {

	return &impl{
		payload:           payload,
		channelComponent:  channelComponent,
		feedbackComponent: feedbackComponent,
	}
}

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*payloadM.PayloadEventResp, error) {

	payloadEventResp := &payloadM.PayloadEventResp{
		ID: im.payload.ID,
	}

	constraintEventResp, err := (*im.channelComponent).Judge(ctx, e)
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

	} else {

		feedReturn, err = im.processFeedReturn(ctx, e, false)

		if err != nil {
			return nil, err
		}

		payloadEventResp.FeedReturn = feedReturn

	}
	return payloadEventResp, nil
}

func (im *impl) processFeedReturn(ctx context.Context, e *eventM.Event, pass bool) (*feedbackM.FeedReturn, error) {
	logrus.Info("payload.processFeedReturn")

	// 計算回饋額
	feedReturn, err := (*im.feedbackComponent).Calculate(ctx, e, pass)
	if err != nil {
		return nil, err
	}

	logrus.Info("payload.processFeedReturn : ", feedReturn)
	return feedReturn, nil
}
