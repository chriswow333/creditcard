package reward_channel

import (
	"context"

	"example.com/creditcard/models/reward_channel"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type impl struct {
	dig.In

	psql *pgx.ConnPool
}

func New(psql *pgx.ConnPool) Store {
	return &impl{
		psql: psql,
	}
}

const INSERT_REWARD_STAT = "INSERT INTO reward_channel " +
	"(\"id\", \"order\", \"card_id\", \"card_reward_id\", \"channel_id\", \"channel_type\") " +
	" VALUES($1, $2, $3, $4, $5, $6)"

func (im *impl) Create(ctx context.Context, rewardChannel *reward_channel.RewardChannel) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		rewardChannel.ID,
		rewardChannel.Order,
		rewardChannel.CardID,
		rewardChannel.CardRewardID,
		rewardChannel.ChannelID,
		rewardChannel.ChannelType,
	}
	if _, err := tx.Exec(INSERT_REWARD_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		})
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_BY_REWARD_ID_STAT = "SELECT \"id\", \"order\", card_id, card_reward_id, \"channel_id\", \"channel_type\" " +
	" FROM reward_channel WHERE \"card_reward_id\" = $1"

func (im *impl) GetByRewardID(ctx context.Context, cardRewardID string) ([]*reward_channel.RewardChannel, error) {

	rewardChannels := []*reward_channel.RewardChannel{}
	conditions := []interface{}{
		cardRewardID,
	}

	rows, err := im.psql.Query(SELECT_BY_REWARD_ID_STAT, conditions...)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		rewardChannel := &reward_channel.RewardChannel{}

		selector := []interface{}{
			&rewardChannel.ID,
			&rewardChannel.Order,
			&rewardChannel.CardID,
			&rewardChannel.CardRewardID,
			&rewardChannel.ChannelID,
			&rewardChannel.ChannelType,
		}
		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)

			return nil, err
		}

		rewardChannels = append(rewardChannels, rewardChannel)

	}

	return rewardChannels, nil
}
