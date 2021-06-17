package constraintpayload

import (
	"context"

	"example.com/creditcard/components/constraint"
	eventM "example.com/creditcard/models/event"

	constraintM "example.com/creditcard/models/constraint"
)

type impl struct {
	constraints  []*constraint.Component
	operatorType constraintM.OperatorType
}

func New(
	constraints []*constraint.Component,
	operatorType constraintM.OperatorType,
) constraint.Component {

	return &impl{
		constraints:  constraints,
		operatorType: operatorType,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Constraint, error) {

	resps := []*eventM.Response{}

	for _, c := range im.constraints {
		resp, err := (*c).Judge(ctx, e)
		if err != nil {
			return nil, err
		}
		resps = append(resps, resp)
	}

	return nil, nil

}
