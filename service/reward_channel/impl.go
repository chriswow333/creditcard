package reward_channel

import (
	"context"
	"runtime/debug"

	rewardChannelM "example.com/creditcard/models/reward_channel"
	"example.com/creditcard/stores/reward_channel"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type impl struct {
	dig.In

	rewardChannelStore reward_channel.Store
}

func New(
	rewardChannelStore reward_channel.Store,
) Service {
	return &impl{
		rewardChannelStore: rewardChannelStore,
	}
}

func (im *impl) Create(ctx context.Context, rewardChannels []*rewardChannelM.RewardChannel) error {

	for _, rc := range rewardChannels {
		id, err := uuid.NewV4()
		if err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return err
		}
		rc.ID = id.String()
		if err := im.rewardChannelStore.Create(ctx, rc); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return err
		}
	}

	return nil
}

func (im *impl) GetByRewardID(ctx context.Context, cardRewardID string) ([]*rewardChannelM.RewardChannel, error) {

	rewardConstraints, err := im.rewardChannelStore.GetByRewardID(ctx, cardRewardID)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}
	return rewardConstraints, nil
}
