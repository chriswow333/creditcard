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
		return reward, nil
	}

	constraint, err := (*im.constraintComps).Judge(ctx, e)
	if err != nil {
		return nil, err
	}

	reward.Constraint = constraint
	reward.Pass = constraint.Pass
	return reward, nil
}
