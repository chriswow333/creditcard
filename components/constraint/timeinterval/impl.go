package timeinterval

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
	"example.com/creditcard/models/timeinterval"
)

type impl struct {
	timeIntervals []*timeinterval.TimeInterval
	operator      constraintM.OperatorType
	name          string
	descs         []string
}

func New(
	constraintPayload *constraintM.ConstraintPayload,
) constraint.Component {
	return &impl{
		timeIntervals: constraintPayload.TimeIntervals,
		operator:      constraintPayload.Operator,
		name:          constraintPayload.Name,
		descs:         constraintPayload.Descs,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Constraint, error) {

	// TODO Get Range from time
	constraint := &eventM.Constraint{
		Name:           im.name,
		Descs:          im.descs,
		ConstraintType: constraintM.TimeIntervalType,
	}

	//Using calandar

	return constraint, nil
}
