package reward

import (
	"context"

	constraintComp "example.com/creditcard/components/constraint"
	costComp "example.com/creditcard/components/cost"

	constraintM "example.com/creditcard/models/constraint"
	eventM "example.com/creditcard/models/event"
	rewardM "example.com/creditcard/models/reward"
)

type impl struct {
	reward *rewardM.Reward

	costComp        *costComp.Component
	constraintComps []*constraintComp.Component
}

func New(
	reward *rewardM.Reward,
	costComp *costComp.Component,
	constraintComps []*constraintComp.Component,
) Component {

	return &impl{
		reward:          reward,
		costComp:        costComp,
		constraintComps: constraintComps,
	}
}

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*eventM.RewardResp, error) {

	reward := &eventM.RewardResp{
		Name: im.reward.Name,
		Desc: im.reward.Desc,
		Cost: im.reward.Cost,
	}

	constraints := []*eventM.ConstraintResp{}

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

	// 計算回饋額
	cost, err := (*im.costComp).Calculate(ctx, e, reward.Pass)
	if err != nil {
		return nil, err
	}
	reward.Cost = cost

	return reward, nil
}
