package streaming

import (
	"context"

	"example.com/creditcard/components/constraint"
	eventM "example.com/creditcard/models/event"
	streamingM "example.com/creditcard/models/streaming"
)

type impl struct {
	streaming *streamingM.Streaming
}

func New(
	streaming *streamingM.Streaming,
) constraint.Component {
	return &impl{
		streaming: streaming,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	resp := &eventM.Response{
		Pass: false,
	}

	for _, streaming := range e.Streamings {
		if streaming.ID == im.streaming.ID {

			resp.Pass = true

			return resp, nil
		}
	}

	return resp, nil
}
