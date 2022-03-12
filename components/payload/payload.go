package payload

import (
	"context"

	eventM "example.com/creditcard/models/event"

	payloadM "example.com/creditcard/models/payload"
)

type Component interface {
	Satisfy(ctx context.Context, e *eventM.Event) (*payloadM.PayloadResp, error)
}
