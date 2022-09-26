package reward_channel

import (
	"context"

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
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return err
		}
		rc.ID = id.String()
		if err := im.rewardChannelStore.Create(ctx, rc); err != nil {
			logrus.Error(err)
			return err
		}
	}

	return nil
}

func (im *impl) GetByRewardID(ctx context.Context, cardRewardID string) ([]*rewardChannelM.RewardChannel, error) {

	rewardConstraints, err := im.rewardChannelStore.GetByRewardID(ctx, cardRewardID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return rewardConstraints, nil
}
