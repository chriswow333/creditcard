package supermarket

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
	"example.com/creditcard/models/supermarket"
	supermarketM "example.com/creditcard/models/supermarket"
)

type impl struct {
	supermarkets   []*supermarketM.Supermarket
	operator       constraintM.OperatorType
	constraintType constraintM.ConstraintType
	name           string
	descs          []string
}

func New(
	constraintPayload *constraintM.ConstraintPayload,
) constraint.Component {
	return &impl{
		supermarkets:   constraintPayload.Supermarkets,
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
	supermarketMap := make(map[string]*supermarket.Supermarket)

	for _, su := range e.Supermarkets {
		supermarketMap[su.ID] = su
	}

	for _, ec := range im.supermarkets {
		if _, ok := supermarketMap[ec.ID]; ok {
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
