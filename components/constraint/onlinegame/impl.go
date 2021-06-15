package onlinegame

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
	onlinegameM "example.com/creditcard/models/onlinegame"
)

type impl struct {
	onlinegames []*onlinegameM.Onlinegame
	operator    constraintM.OperatorType
}

func New(
	onlinegames []*onlinegameM.Onlinegame,
	operator constraintM.OperatorType,

) constraint.Component {

	return &impl{
		onlinegames: onlinegames,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	resp := &eventM.Response{
		Pass: false,
	}

	return resp, nil
}
