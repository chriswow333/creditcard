package onlinegame

import (
	"context"

	"example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
	onlinegameM "example.com/creditcard/models/onlinegame"
)

type impl struct {
	onlinegames        []*onlinegameM.Onlinegame
	constraintType     constraintM.ConstraintType
	name               string
	constraintOperator constraintM.OperatorType
}

func New(
	constraint *constraintM.Constraint,
) constraint.Component {

	return &impl{
		onlinegames:        constraint.Onlinegames,
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
	onlinegameMap := make(map[string]*onlinegameM.Onlinegame)

	for _, on := range e.Onlinegames {
		onlinegameMap[on.ID] = on
	}

	for _, on := range im.onlinegames {
		if _, ok := onlinegameMap[on.ID]; ok {
			matches = append(matches, on.ID)
		} else {
			misses = append(misses, on.ID)
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
