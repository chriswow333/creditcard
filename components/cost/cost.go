package cost

import (
	"context"

	costM "example.com/creditcard/models/cost"
	eventM "example.com/creditcard/models/event"
)

type Component interface {
	Calculate(ctx context.Context, e *eventM.Event, pass bool) (*costM.Cost, error)
}
