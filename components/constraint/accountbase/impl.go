package accountbase

import (
	"context"

	"example.com/creditcard/components/constraint"
	baseM "example.com/creditcard/models/base"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	accountBases []*baseM.AccountBase
	operator     constraintM.OperatorType
}

func New(
	accountBases []*baseM.AccountBase,
	operator constraintM.OperatorType,
) constraint.Component {

	return &impl{
		accountBases: accountBases,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	resp := &eventM.Response{
		Pass: true,
	}

	return resp, nil
}
