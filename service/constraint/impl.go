package constraint

import (
	"context"

	constraintM "example.com/creditcard/models/constraint"
	"example.com/creditcard/stores/reward"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	uuid "github.com/nu7hatch/gouuid"
)

type impl struct {
	dig.In

	rewardStore reward.Store
}

func New(
	rewardStore reward.Store,
) Service {
	return &impl{
		rewardStore: rewardStore,
	}
}

func (im *impl) UpdateByRewardID(ctx context.Context, rewardID string, constraints []*constraintM.Constraint) error {

	reward, err := im.rewardStore.GetByID(ctx, rewardID)
	if err != nil {
		logrus.Error(err)
		return err
	}

	for _, c := range constraints {
		if c.ID != "" {
			id, err := uuid.NewV4()
			if err != nil {
				logrus.Error(err)
				return err
			}
			c.ID = id.String()
		}
	}

	reward.Constraints = constraints
	if err := im.rewardStore.UpdateByID(ctx, reward); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) GetByRewardID(ctx context.Context, rewardID string) ([]*constraintM.Constraint, error) {
	rewardModel, err := im.rewardStore.GetByID(ctx, rewardID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return rewardModel.Constraints, nil
}
