package timeinterval

import (
	"context"
	"time"

	"example.com/creditcard/components/constraint"

	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
	timeintervalM "example.com/creditcard/models/timeinterval"
)

type impl struct {
	timeIntervals []*timeintervalM.TimeInterval
	operator      constraintM.OperatorType
	name          string
	desc          string
}

func New(
	constraintPayload *constraintM.ConstraintPayload,
) constraint.Component {
	impl := &impl{
		timeIntervals: constraintPayload.TimeIntervals,
		operator:      constraintPayload.Operator,
		name:          constraintPayload.Name,
		desc:          constraintPayload.Desc,
	}

	return impl
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.ConstraintResp, error) {

	// TODO Get Range from time
	constraint := &eventM.ConstraintResp{
		Name:           im.name,
		Desc:           im.desc,
		ConstraintType: constraintM.TimeIntervalType,
	}

	matches := []string{}
	misses := []string{}

	for _, t := range im.timeIntervals {
		switch t.TimeType {
		case timeintervalM.WeekDay:
			weekDay := time.Unix(e.EffictiveTime, 0).Weekday()
			if t.WeekDayFrom <= int32(weekDay) && int32(weekDay) <= t.WeekDayTo {

				matches = append(matches, t.ID)
			} else {
				misses = append(misses, t.ID)
			}
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
