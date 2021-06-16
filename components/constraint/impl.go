package constraint

import (
	"context"

	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	constraint *constraintM.Constraint
	component  *Component
}

func New(
	constraint *constraintM.Constraint,
	component *Component,
) Component {

	return &impl{
		constraint: constraint,
		component:  component,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	resp, err := (*im.component).Judge(ctx, e)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
