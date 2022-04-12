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
	constraintResp *constraintM.ConstraintResp
}

func New(
	constraintResp *constraintM.ConstraintResp,
) constraint.Component {

	impl := &impl{
		constraintResp: constraintResp,
	}

	return impl
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*constraintM.ConstraintEventResp, error) {

	// TODO Get Range from time

	constraintEventResp := &constraintM.ConstraintEventResp{}

	matches := []string{}
	misses := []string{}

	for _, t := range im.constraintResp.TimeIntervals {
		switch t.TimeType {
		case timeintervalM.WeekDay:
			weekDay := time.Unix(e.EffectiveTime, 0).Weekday()
			if t.WeekDayFrom <= int32(weekDay) && int32(weekDay) <= t.WeekDayTo {
				matches = append(matches, t.ID)
			} else {
				misses = append(misses, t.ID)
			}
		}
	}

	constraintEventResp.Matches = matches
	constraintEventResp.Misses = misses
	switch im.constraintResp.ConstraintOperatorType {
	case constraintM.OR:
		if len(matches) > 0 {
			constraintEventResp.Pass = true
		} else {
			constraintEventResp.Pass = false
		}
	case constraintM.AND:
		if len(misses) > 0 {
			constraintEventResp.Pass = false
		} else {
			constraintEventResp.Pass = true
		}
	}

	return constraintEventResp, nil
}
