package accountbase

import (
	"context"

	"example.com/creditcard/components/constraint"
	baseM "example.com/creditcard/models/base"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	accountBases   []*baseM.AccountBase
	operator       constraintM.OperatorType
	constraintType constraintM.ConstraintType
	name           string
	descs          []string
}

func New(
	constraintPayload *constraintM.ConstraintPayload,
) constraint.Component {

	return &impl{
		accountBases:   constraintPayload.AccountBases,
		operator:       constraintPayload.Operator,
		constraintType: constraintPayload.ConstraintType,
		name:           constraintPayload.Name,
		descs:          constraintPayload.Descs,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Constraint, error) {

	resp := &eventM.Response{
		Pass: true,
	}

	return resp, nil
}
