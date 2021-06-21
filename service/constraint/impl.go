package constraint

import (
	"context"

	constraintM "example.com/creditcard/models/constraint"
	"example.com/creditcard/stores/reward"
)

type impl struct {
	rewardStore reward.Store
}

func New(
	rewardStore reward.Store,
) Service {
	return &impl{
		rewardStore: rewardStore,
	}
}

func (im *impl) Create(ctx context.Context, rewardID string, constraints []*constraintM.Constraint) error {
	reward, err := im.rewardStore.GetByID(ctx, rewardID)
	if err != nil {
		return err
	}
	reward.Constraints = constraints
	if err := im.rewardStore.UpdateByID(ctx, reward); err != nil {
		return err
	}
	return nil
}

func (im *impl) GetByRewardID(ctx context.Context, rewardID string) ([]*constraintM.Constraint, error) {
	rewardModel, err := im.rewardStore.GetByID(ctx, rewardID)
	if err != nil {
		return nil, err
	}
	return rewardModel.Constraints, nil
}
