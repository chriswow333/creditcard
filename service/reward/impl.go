package reward

import (
	"context"
	"time"

	rewardM "example.com/creditcard/models/reward"
	"example.com/creditcard/stores/reward"
	"github.com/sirupsen/logrus"

	"go.uber.org/dig"

	uuid "github.com/nu7hatch/gouuid"
)

var (
	timeNow = time.Now
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

func (im *impl) Create(ctx context.Context, reward *rewardM.Reward) error {

	reward.UpdateDate = timeNow().Unix()

	id, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	reward.ID = id.String()

	if err := im.rewardStore.Create(ctx, reward); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (im *impl) GetByID(ctx context.Context, ID string) (*rewardM.Reward, error) {
	reward, err := im.rewardStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return reward, nil
}

func (im *impl) GetByCardID(ctx context.Context, cardID string) ([]*rewardM.Reward, error) {
	rewards, err := im.rewardStore.GetByCardID(ctx, cardID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return rewards, nil
}

func (im *impl) UpdateByID(ctx context.Context, reward *rewardM.Reward) error {
	if err := im.rewardStore.UpdateByID(ctx, reward); err != nil {
		logrus.Error(err)
		return err
	}
	return nil

}
