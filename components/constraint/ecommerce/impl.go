package ecommerce

import (
	"context"

	"example.com/creditcard/components/constraint"
	ecommerceM "example.com/creditcard/models/ecommerce"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	ecommerce *ecommerceM.Ecommerce
}

func New(
	ecommerce *ecommerceM.Ecommerce,

) constraint.Component {

	return &impl{
		ecommerce: ecommerce,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	resp := &eventM.Response{
		Pass: false,
	}

	for _, ec := range e.Ecommerces {
		if ec.ID == im.ecommerce.ID {
			resp.Pass = true
			return resp, nil
		}
	}
	return resp, nil
}
