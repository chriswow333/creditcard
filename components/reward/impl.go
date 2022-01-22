package reward

import (
	"context"

	constraintComp "example.com/creditcard/components/constraint"

	eventM "example.com/creditcard/models/event"
	rewardM "example.com/creditcard/models/reward"
)

type impl struct {
	reward *rewardM.Reward

	constraintComps *constraintComp.Component
}

func New(
	reward *rewardM.Reward,
	constraintComps *constraintComp.Component,
) Component {

	return &impl{
		reward:          reward,
		constraintComps: constraintComps,
	}
}

func (im *impl) Satisfy(ctx context.Context, e *eventM.Event) (*eventM.RewardResp, error) {

	reward := &eventM.RewardResp{
		Name: im.reward.Name,
		Desc: im.reward.Desc,
	}

	if !(im.reward.StartDate <= e.EffictiveTime && e.EffictiveTime <= im.reward.EndDate) {
		// cost, err := im.processCost(ctx, e, false)
		// if err != nil {
		// 	return nil, err
		// }
		// reward.Cost = cost
		return reward, nil
	}

	constraint, err := (*im.constraintComps).Judge(ctx, e)
	if err != nil {
		return nil, err
	}

	reward.Constraint = constraint
	reward.Pass = constraint.Pass
	return reward, nil
	// for _, c := range im.constraintComps {

	// 	constraint, err := (*c).Judge(ctx, e)

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	if constraint.Pass {
	// 		matches++
	// 	} else {
	// 		misses++
	// 	}

	// 	constraints = append(constraints, constraint)
	// }

	// reward.Constraints = constraints

	// switch im.reward.Operator {
	// case constraintM.OrOperator:
	// 	if matches > 0 {
	// 		reward.Pass = true
	// 	} else {
	// 		reward.Pass = false
	// 	}
	// case constraintM.AndOperator:
	// 	if misses > 0 {
	// 		reward.Pass = false
	// 	} else {
	// 		reward.Pass = true
	// 	}

	// }

	// cost, err := im.processCost(ctx, e, reward.Pass)

	// if err != nil {
	// 	return nil, err
	// }

	// // without bonus is fail even if the constraint is pass.
	// if !cost.IsRewardGet {
	// 	reward.Pass = false
	// }

	// reward.Cost = cost

	// return reward, nil
}
