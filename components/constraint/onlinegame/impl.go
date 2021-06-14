package onlinegame

import (
	"context"

	"example.com/creditcard/components/constraint"
	eventM "example.com/creditcard/models/event"
	onlinegameM "example.com/creditcard/models/onlinegame"
)

type impl struct {
	onlinegame *onlinegameM.Onlinegame
}

func New(
	onlinegame *onlinegameM.Onlinegame,
) constraint.Component {

	return &impl{
		onlinegame: onlinegame,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	resp := &eventM.Response{
		Pass: false,
	}

	for _, online := range e.Onlinegames {
		if online.ID == im.onlinegame.ID {
			resp.Pass = true
			return resp, nil
		}
	}

	return resp, nil
}
