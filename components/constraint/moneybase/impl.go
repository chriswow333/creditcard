package moneybase

import (
	"context"

	"example.com/creditcard/components/constraint"
	baseM "example.com/creditcard/models/base"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	moneybase *baseM.MoneyBase
}

func New(
	moneybase *baseM.MoneyBase,

) constraint.Component {

	return &impl{
		moneybase: moneybase,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	return nil, nil
}
