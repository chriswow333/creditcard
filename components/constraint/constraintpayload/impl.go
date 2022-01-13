package constraintpayload

import (
	"context"

	"example.com/creditcard/components/constraint"
	eventM "example.com/creditcard/models/event"

	constraintM "example.com/creditcard/models/constraint"
)

type impl struct {
	constraints []*constraint.Component
	operator    constraintM.OperatorType
	name        string
	desc        string
}

func New(
	constraints []*constraint.Component,
	constraintPayload *constraintM.ConstraintPayload,
) constraint.Component {

	return &impl{
		constraints: constraints,
		operator:    constraintPayload.Operator,
		name:        constraintPayload.Name,
		desc:        constraintPayload.Desc,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.ConstraintResp, error) {

	constraint := &eventM.ConstraintResp{
		Name:           im.name,
		Desc:           im.desc,
		ConstraintType: constraintM.ConstraintPayloadType,
	}

	eventConstraints := []*eventM.ConstraintResp{}

	matches := 0
	misses := 0
	for _, c := range im.constraints {
		constraintResp, err := (*c).Judge(ctx, e)
		if err != nil {
			return nil, err
		}
		eventConstraints = append(eventConstraints, constraintResp)
		if constraintResp.Pass {
			matches++
		} else {
			misses++
		}
	}

	switch im.operator {
	case constraintM.OrOperator:
		if matches > 0 {
			constraint.Pass = true
		} else {
			constraint.Pass = false
		}
	case constraintM.AndOperator:
		if misses > 0 {
			constraint.Pass = false
		} else {
			constraint.Pass = true
		}
	}

	constraint.Constraints = eventConstraints

	return constraint, nil

}
