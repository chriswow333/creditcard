package reward

import (
	"context"

	constraintComp "example.com/creditcard/components/constraint"
	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
	rewardM "example.com/creditcard/models/reward"
)

type impl struct {
	reward          *rewardM.Reward
	constraintComps []*constraintComp.Component
}

func New(
	reward *rewardM.Reward,
	constraintComps []*constraintComp.Component,
) Component {

	return &impl{
		reward:          reward,
		constraintComps: constraintComps,
	}
}

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*eventM.Reward, error) {

	reward := &eventM.Reward{
		Name:  im.reward.Name,
		Desc:  im.reward.Desc,
		Bonus: im.reward.Bonus,
	}

	constraints := []*eventM.Constraint{}

	matches := 0
	misses := 0

	for _, c := range im.constraintComps {
		constraint, err := (*c).Judge(ctx, e)
		if err != nil {
			return nil, err
		}

		if constraint.Pass {
			matches++
		} else {
			misses++
		}
		constraints = append(constraints, constraint)
	}

	reward.Constraints = constraints
	switch im.reward.Operator {
	case constraintM.OrOperator:
		if matches > 0 {
			reward.Pass = true
		} else {
			reward.Pass = false
		}
	case constraintM.AndOperator:
		if misses > 0 {
			reward.Pass = false
		} else {
			reward.Pass = true
		}
	}

	return reward, nil
}
