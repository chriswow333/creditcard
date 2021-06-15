package moneybase

import (
	"context"

	"example.com/creditcard/components/constraint"
	baseM "example.com/creditcard/models/base"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	moneybases []*baseM.MoneyBase
	operator   constraintM.OperatorType
}

func New(
	moneybases []*baseM.MoneyBase,
	operator constraintM.OperatorType,
) constraint.Component {

	return &impl{
		moneybases: moneybases,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	return nil, nil
}
