package time

import (
	"context"

	"example.com/creditcard/components/constraint"
	"example.com/creditcard/models/base"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	timeBases []*base.TimeBase
	operator  constraintM.OperatorType
}

func New(
	timeBases []*base.TimeBase,
	operator constraintM.OperatorType,
) constraint.Component {

	return &impl{
		timeBases: timeBases,
	}

}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	// TODO Get Range from time

	return nil, nil
}
