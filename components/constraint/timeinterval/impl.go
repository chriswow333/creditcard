package time

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
	"example.com/creditcard/models/timeinterval"
)

type impl struct {
	timeIntervals  []*timeinterval.TimeInterval
	operator       constraintM.OperatorType
	constratinType constraintM.ConstraintType
	name           string
	descs          []string
}

func New(
	constraintPayload *constraintM.ConstraintPayload,
) constraint.Component {
	return &impl{
		timeIntervals:  constraintPayload.TimeIntervals,
		operator:       constraintPayload.Operator,
		constratinType: constraintPayload.ConstraintType,
		name:           constraintPayload.Name,
		descs:          constraintPayload.Descs,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Constraint, error) {

	// TODO Get Range from time

	return nil, nil
}
