package feedback

import (
	"context"

	eventM "example.com/creditcard/models/event"
	feedbackM "example.com/creditcard/models/feedback"
)

type Component interface {
	Calculate(ctx context.Context, e *eventM.Event, pass bool) (*feedbackM.Feedback, error)
}
