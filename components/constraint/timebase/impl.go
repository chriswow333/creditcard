package time

import (
	"context"

	"example.com/creditcard/components/constraint"
	"example.com/creditcard/models/base"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	timeBase *base.TimeBase
}

func New(
	timeBase *base.TimeBase,
) constraint.Component {

	return &impl{
		timeBase: timeBase,
	}

}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	// TODO Get Range from time

	return nil, nil
}
