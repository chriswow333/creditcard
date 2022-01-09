package onlinegame

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
	"example.com/creditcard/models/onlinegame"
	onlinegameM "example.com/creditcard/models/onlinegame"
)

type impl struct {
	onlinegames    []*onlinegameM.Onlinegame
	constraintType constraintM.ConstraintType
	name           string
	descs          []string
	operator       constraintM.OperatorType
}

func New(
	constraintPayload *constraintM.ConstraintPayload,
) constraint.Component {

	return &impl{
		onlinegames:    constraintPayload.Onlinegames,
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
	onlinegameMap := make(map[string]*onlinegame.Onlinegame)

	for _, on := range e.Onlinegames {
		onlinegameMap[on.ID] = on
	}

	for _, ec := range im.onlinegames {
		if _, ok := onlinegameMap[ec.ID]; ok {
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
