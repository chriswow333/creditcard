package reward

import (
	"context"
	"runtime/debug"
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

	id, err := uuid.NewV4()
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	reward.ID = id.String()

	if err := im.rewardStore.Create(ctx, reward); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	return nil
}

func (im *impl) GetByID(ctx context.Context, ID string) (*rewardM.Reward, error) {
	reward, err := im.rewardStore.GetByID(ctx, ID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return reward, nil
}

func (im *impl) GetByCardID(ctx context.Context, cardID string) ([]*rewardM.Reward, error) {
	rewards, err := im.rewardStore.GetByCardRewardID(ctx, cardID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}
	return rewards, nil
}

func (im *impl) UpdateByID(ctx context.Context, reward *rewardM.Reward) error {
	if err := im.rewardStore.UpdateByID(ctx, reward); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	return nil

}
