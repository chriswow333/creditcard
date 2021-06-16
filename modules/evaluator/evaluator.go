package evaluator

import (
	"context"

	eventM "example.com/creditcard/models/event"
)

type Module interface {
	Evaluate(ctx context.Context, e *eventM.Event) (*eventM.Response, error)
}
