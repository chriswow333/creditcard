package constraint

import (
	"context"

	eventM "example.com/creditcard/models/event"
)

type Component interface {
	Judge(ctx context.Context, e *eventM.Event) (*eventM.ConstraintResp, error)
}
