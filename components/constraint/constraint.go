package constraint

import (
	"context"

	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
)

type Component interface {
	Judge(ctx context.Context, e *eventM.Event) (*constraintM.ConstraintResp, error)
}
