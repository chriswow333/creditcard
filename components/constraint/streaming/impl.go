package streaming

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
	streamingM "example.com/creditcard/models/streaming"
)

type impl struct {
	streamings []*streamingM.Streaming
	operator   constraintM.OperatorType
}

func New(
	streamings []*streamingM.Streaming,
	operator constraintM.OperatorType,
) constraint.Component {
	return &impl{
		streamings: streamings,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	resp := &eventM.Response{
		Pass: false,
	}

	return resp, nil
}
