package mobilepay

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
	"example.com/creditcard/models/mobilepay"
	mobilepayM "example.com/creditcard/models/mobilepay"
)

type impl struct {
	mobilepays     []*mobilepayM.Mobilepay
	operator       constraintM.OperatorType
	constraintType constraintM.ConstraintType
	name           string
	descs          []string
}

func New(
	constraintPayload *constraintM.ConstraintPayload,
) constraint.Component {
	return &impl{
		mobilepays:     constraintPayload.Mobilepays,
		operator:       constraintPayload.Operator,
		constraintType: constraintPayload.ConstraintType,
		name:           constraintPayload.Name,
		descs:          constraintPayload.Descs,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.ConstraintResp, error) {

	constraint := &eventM.ConstraintResp{
		Name:           im.name,
		Descs:          im.descs,
		ConstraintType: im.constraintType,
	}

	matches := []string{}
	misses := []string{}

	mobilepayMap := make(map[string]*mobilepay.Mobilepay)

	for _, mo := range e.Mobilepays {
		mobilepayMap[mo.ID] = mo

	}

	for _, mo := range im.mobilepays {

		if _, ok := mobilepayMap[mo.ID]; ok {
			matches = append(matches, mo.ID)
		} else {
			misses = append(misses, mo.ID)
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
