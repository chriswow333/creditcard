package mobilepay

import (
	"context"

	"example.com/creditcard/components/constraint"
	eventM "example.com/creditcard/models/event"
	mobilepayM "example.com/creditcard/models/mobilepay"
)

type impl struct {
	mobilepay *mobilepayM.Mobilepay
}

func New(
	mobilepay *mobilepayM.Mobilepay,
) constraint.Component {
	return &impl{
		mobilepay: mobilepay,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	resp := &eventM.Response{
		Pass: false,
	}

	for _, pay := range e.Mobilepays {

		if pay.ID == im.mobilepay.ID {
			resp.Pass = true
			return resp, nil
		}
	}

	return resp, nil
}
