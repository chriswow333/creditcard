package card

import (
	"context"

	eventM "example.com/creditcard/models/event"

	cardM "example.com/creditcard/models/card"
)

type Component interface {
	Satisfy(ctx context.Context, e *eventM.Event) (*cardM.CardResp, error)
}
