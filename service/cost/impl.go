package cost

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	costM "example.com/creditcard/models/cost"
	"example.com/creditcard/stores/reward"
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

func (im *impl) UpdateByRewardID(ctx context.Context, rewardID string, cost *costM.Cost) error {

	reward, err := im.rewardStore.GetByID(ctx, rewardID)
	if err != nil {
		logrus.Error(err)
		return err
	}

	reward.Cost = cost

	if err := im.rewardStore.UpdateByID(ctx, reward); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) GetByRewardID(ctx context.Context, rewardID string) (*costM.Cost, error) {

	rewardModel, err := im.rewardStore.GetByID(ctx, rewardID)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return rewardModel.Cost, nil
}
