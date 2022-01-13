package supermarket

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
	supermarketM "example.com/creditcard/models/supermarket"
)

type impl struct {
	supermarkets   []*supermarketM.Supermarket
	operator       constraintM.OperatorType
	constraintType constraintM.ConstraintType
	name           string
	desc           string
}

func New(
	constraintPayload *constraintM.ConstraintPayload,
) constraint.Component {
	return &impl{
		supermarkets:   constraintPayload.Supermarkets,
		operator:       constraintPayload.Operator,
		constraintType: constraintPayload.ConstraintType,
		name:           constraintPayload.Name,
		desc:           constraintPayload.Desc,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.ConstraintResp, error) {

	constraint := &eventM.ConstraintResp{
		Name:           im.name,
		Desc:           im.desc,
		ConstraintType: im.constraintType,
	}

	matches := []string{}
	misses := []string{}
	supermarketMap := make(map[string]*supermarketM.Supermarket)

	for _, su := range e.Supermarkets {
		supermarketMap[su.ID] = su
	}

	for _, su := range im.supermarkets {
		if _, ok := supermarketMap[su.ID]; ok {
			matches = append(matches, su.ID)
		} else {
			misses = append(misses, su.ID)
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
