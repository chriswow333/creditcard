package constraintpayload

import (
	"context"

	"example.com/creditcard/components/constraint"
	costComp "example.com/creditcard/components/cost"

	costM "example.com/creditcard/models/cost"
	eventM "example.com/creditcard/models/event"

	constraintM "example.com/creditcard/models/constraint"
)

type impl struct {
	constraints []*constraint.Component

	costComp *costComp.Component

	operator constraintM.OperatorType
	name     string
	desc     string
}

func New(
	constraints []*constraint.Component,
	costComp *costComp.Component,
	constraintPayload *constraintM.ConstraintPayload,
) constraint.Component {

	return &impl{
		constraints: constraints,
		costComp:    costComp,

		operator: constraintPayload.Operator,
		name:     constraintPayload.Name,
		desc:     constraintPayload.Desc,
	}
}

func (im *impl) Judge(ctx context.Context, e *eventM.Event) (*eventM.ConstraintResp, error) {

	constraint := &eventM.ConstraintResp{
		Name:           im.name,
		Desc:           im.desc,
		ConstraintType: constraintM.ConstraintPayloadType,
	}

	eventConstraints := []*eventM.ConstraintResp{}

	matches := 0
	misses := 0
	for _, c := range im.constraints {
		constraintResp, err := (*c).Judge(ctx, e)
		if err != nil {
			return nil, err
		}
		eventConstraints = append(eventConstraints, constraintResp)
		if constraintResp.Pass {
			matches++
		} else {
			misses++
		}
	}

	switch im.operator {
	case constraintM.OrOperator:
		if matches > 0 {
			constraint.Pass = true
		} else {
			constraint.Pass = false
		}
	case constraintM.AndOperator:
		if misses > 0 {
			constraint.Pass = false
		} else {
			constraint.Pass = true
		}
	}

	constraint.Constraints = eventConstraints

	if im.costComp != nil {

		var cost *costM.Cost
		var err error

		if constraint.Pass {
			cost, err = im.processCost(ctx, e, true)
			if err != nil {
				return nil, err
			}
		} else {
			cost, err = im.processCost(ctx, e, false)

		}
		if err != nil {
			return nil, err
		}

		constraint.Cost = cost

		if !cost.IsRewardGet {
			constraint.Pass = false
		}

	}

	return constraint, nil

}

func (im *impl) processCost(ctx context.Context, e *eventM.Event, pass bool) (*costM.Cost, error) {

	// 計算回饋額
	cost, err := (*im.costComp).Calculate(ctx, e, pass)
	if err != nil {
		return nil, err
	}

	return cost, nil
}
