package ecommerce

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	ecommerceM "example.com/creditcard/models/ecommerce"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	ecommerceMap map[string]*ecommerceM.Ecommerce
	operator     constraintM.OperatorType
}

func New(
	ecommerces []*ecommerceM.Ecommerce,
	operator constraintM.OperatorType,

) constraint.Component {

	m := make(map[string]*ecommerceM.Ecommerce)
	for _, e := range ecommerces {
		if _, ok := m[e.ID]; ok {
			panic("duplicate")
		} else {
			m[e.ID] = e
		}
	}

	return &impl{
		ecommerceMap: m,
		operator:     operator,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Response, error) {

	resp := &eventM.Response{}
	if im.operator == constraintM.OrOperator {
		for _, e := range e.Ecommerces {
			if _, ok := im.ecommerceMap[e.ID]; ok {
				resp.Pass = true
				return resp, nil
			}
		}

		resp.Pass = false
		return resp, nil
	} else {
		for _, e := range e.Ecommerces {
			if _, ok := im.ecommerceMap[e.ID]; !ok {
				resp.Pass = false
				return resp, nil
			}

		}

		resp.Pass = true
		return resp, nil
	}
}
