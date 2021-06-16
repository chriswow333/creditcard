package supermarket

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
	supermarketM "example.com/creditcard/models/supermarket"
)

type impl struct {
	supermarkets []*supermarketM.Supermarket
	operator     constraintM.OperatorType
}

func New(
	supermarkets []*supermarketM.Supermarket,
	operator constraintM.OperatorType,
) constraint.Component {

	return &impl{
		supermarkets: supermarkets,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	resp := &eventM.Response{
		Pass: false,
	}

	return resp, nil
}
