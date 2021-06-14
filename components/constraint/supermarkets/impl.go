package constraint

import (
	"context"

	"example.com/creditcard/components/constraint"
	eventM "example.com/creditcard/models/event"
	supermarketM "example.com/creditcard/models/supermarket"
)

type impl struct {
	supermarket *supermarketM.Supermarket
}

func New(
	supermarket *supermarketM.Supermarket,
) constraint.Component {

	return &impl{
		supermarket: supermarket,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	resp := &eventM.Response{
		Pass: false,
	}

	for _, super := range e.Supermarkets {

		if super.ID == im.supermarket.ID {
			resp.Pass = true
			return resp, nil
		}
	}

	return resp, nil
}
