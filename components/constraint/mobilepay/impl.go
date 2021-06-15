package mobilepay

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
	mobilepayM "example.com/creditcard/models/mobilepay"
)

type impl struct {
	mobilepays []*mobilepayM.Mobilepay
	operator   constraintM.OperatorType
}

func New(
	mobilepays []*mobilepayM.Mobilepay,
	operator constraintM.OperatorType,
) constraint.Component {
	return &impl{
		mobilepays: mobilepays,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	resp := &eventM.Response{
		Pass: false,
	}

	return resp, nil
}
