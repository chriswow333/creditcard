package ecommerce

import (
	"context"

	"example.com/creditcard/components/constraint"

	constraintM "example.com/creditcard/models/constraint"
	ecommerceM "example.com/creditcard/models/ecommerce"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	ecommerces         []*ecommerceM.Ecommerce
	constraintOperator constraintM.OperatorType
	constraintType     constraintM.ConstraintType
}

func New(
	constraint *constraintM.Constraint,
) constraint.Component {

	return &impl{
		ecommerces:         constraint.Ecommerces,
		constraintOperator: constraint.ConstraintOperator,
		constraintType:     constraint.ConstraintType,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*constraintM.ConstraintResp, error) {

	constraint := &constraintM.ConstraintResp{
		ConstraintType: im.constraintType,
	}

	matches := []string{}
	misses := []string{}

	ecommerceMap := make(map[string]*ecommerceM.Ecommerce)

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

	switch im.constraintOperator {
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
