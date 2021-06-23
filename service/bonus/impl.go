package bonus

import (
	"context"

	bonusM "example.com/creditcard/models/bonus"
	"example.com/creditcard/stores/reward"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
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

func (im *impl) UpdateByRewardID(ctx context.Context, rewardID string, bonus *bonusM.Bonus) error {

	reward, err := im.rewardStore.GetByID(ctx, rewardID)
	if err != nil {
		logrus.Error(err)
		return err
	}
	reward.Bonus = bonus
	if err := im.rewardStore.UpdateByID(ctx, reward); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) GetByRewardID(ctx context.Context, rewardID string) (*bonusM.Bonus, error) {
	rewardModel, err := im.rewardStore.GetByID(ctx, rewardID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return rewardModel.Bonus, nil
}
