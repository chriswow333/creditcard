package ecommerce

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	"example.com/creditcard/models/ecommerce"
	ecommerceM "example.com/creditcard/models/ecommerce"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	ecommerces     []*ecommerceM.Ecommerce
	operator       constraintM.OperatorType
	constraintType constraintM.ConstraintType
	name           string
	descs          []string
}

func New(
	constraintPayload *constraintM.ConstraintPayload,
) constraint.Component {

	return &impl{
		ecommerces:     constraintPayload.Ecommerces,
		operator:       constraintPayload.Operator,
		constraintType: constraintPayload.ConstraintType,
		name:           constraintPayload.Name,
		descs:          constraintPayload.Descs,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Constraint, error) {

	constraint := &eventM.Constraint{
		Name:           im.name,
		Descs:          im.descs,
		ConstraintType: im.constraintType,
	}

	matches := []string{}
	misses := []string{}
	ecommerceMap := make(map[string]*ecommerce.Ecommerce)

	for _, ec := range e.Ecommerces {
		ecommerceMap[ec.ID] = ec
	}

	for _, ec := range im.ecommerces {
		if _, ok := ecommerceMap[ec.ID]; ok {
			matches = append(matches, ec.ID)
		} else {
			misses = append(misses, ec.ID)
		}
	}

	constraint.Matches = matches
	constraint.Misses = misses

	switch im.operator {
	case constraintM.OrOperator:
		if len(matches) > 0 {
			constraint.Pass = true
		} else {
			constraint.Pass = false
		}
	case constraintM.AndOperator:
		if len(misses) > 0 {
			constraint.Pass = false
		} else {
			constraint.Pass = true
		}
	}

	return constraint, nil
}
