package cost

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	costM "example.com/creditcard/models/cost"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	costLimit *costM.CostLimit
}

func New(
	costLimit *costM.CostLimit,
) constraint.Component {

	return &impl{
		costLimit: costLimit,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Constraint, error) {

	constraint := &eventM.Constraint{
		Name:           im.costLimit.Name,
		Descs:          im.costLimit.Descs,
		ConstraintType: constraintM.CostType,
	}

	if e.Cost.Currency == im.costLimit.Currency {
		if e.Cost.Current >= im.costLimit.AtLeast &&
			e.Cost.Current <= im.costLimit.AtMost {
			constraint.Pass = true
		} else {
			constraint.Pass = false
		}
	} else {
		// if empty, set true
		constraint.Pass = true
	}

	return constraint, nil
}
