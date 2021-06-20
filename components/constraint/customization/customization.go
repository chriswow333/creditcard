package customization

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	customizationM "example.com/creditcard/models/customization"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	customizations []*customizationM.Customization
	operator       constraintM.OperatorType
	name           string
	descs          []string
}

func New(
	constraintPayload constraintM.ConstraintPayload,
) constraint.Component {
	return &impl{
		customizations: constraintPayload.Customizations,
		operator:       constraintPayload.Operator,
		name:           constraintPayload.Name,
		descs:          constraintPayload.Descs,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Constraint, error) {

	constraint := &eventM.Constraint{
		Name:           im.name,
		Descs:          im.descs,
		ConstraintType: constraintM.CustomizationType,
	}

	matches := []string{}
	misses := []string{}

	customizationMap := make(map[string]*customizationM.Customization)

	for _, cust := range e.Customizations {
		customizationMap[cust.ID] = cust
	}

	for _, cust := range im.customizations {
		if _, ok := customizationMap[cust.ID]; ok {
			matches = append(matches, cust.ID)
		} else {
			misses = append(misses, cust.ID)
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
