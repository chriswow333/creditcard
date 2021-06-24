package bonuslimit

import (
	"context"

	"example.com/creditcard/components/constraint"
	bonusM "example.com/creditcard/models/bonus"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
)

type impl struct {
	bonusLimit *bonusM.BonusLimit
}

func New(
	constraintPayload *constraintM.ConstraintPayload,
) constraint.Component {
	return &impl{
		bonusLimit: constraintPayload.BonusLimit,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.Constraint, error) {

	constraint := &eventM.Constraint{
		Name:           im.bonusLimit.Name,
		Descs:          []string{im.bonusLimit.Desc},
		ConstraintType: constraintM.BonusLimitType,
	}
	if e.Bonus.BonusType == im.bonusLimit.BonusType {
		// TODO
	} else {
		constraint.Pass = true
	}

	return constraint, nil
}
