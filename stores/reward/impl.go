package reward

import (
	"context"

	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	rewardM "example.com/creditcard/models/reward"
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

const INSERT_REWARD_STAT = "INSERT INTO reward " +
	"(\"id\", \"card_reward_id\", \"order\", \"title\", \"sub_title\", " +
	" start_date, end_date, update_date, payload_operator, payload)" +
	" VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

func (im *impl) Create(ctx context.Context, reward *rewardM.Reward) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		reward.ID,
		reward.CardRewardID,
		reward.Order,
		reward.Title,
		reward.SubTitle,
		reward.StartDate,
		reward.EndDate,
		reward.UpdateDate,
		reward.PayloadOperator,
		reward.Payloads,
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

const SELECT_STAT = "SELECT \"id\", card_reward_id, \"order\", \"title\", \"sub_title\", " +
	" start_date, end_date, update_date, payload_operator, payload " +
	" FROM reward WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*rewardM.Reward, error) {

	reward := &rewardM.Reward{}

	selector := []interface{}{
		&reward.ID,
		&reward.CardRewardID,
		&reward.Order,
		&reward.Title,
		&reward.SubTitle,
		&reward.StartDate,
		&reward.EndDate,
		&reward.UpdateDate,
		&reward.PayloadOperator,
		&reward.Payloads,
	}

	if err := im.psql.QueryRow(SELECT_STAT, ID).Scan(selector...); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return reward, nil
}

const SELECT_BY_CARDID_STAT = "SELECT \"id\", card_reward_id, \"order\", \"title\", \"sub_title\", " +
	" start_date, end_date, update_date, payload_operator, payload " +
	" FROM reward WHERE card_reward_id = $1"

func (im *impl) GetByCardRewardID(ctx context.Context, cardID string) ([]*rewardM.Reward, error) {

	rewards := []*rewardM.Reward{}

	conditions := []interface{}{
		cardID,
	}
	rows, err := im.psql.Query(SELECT_BY_CARDID_STAT, conditions...)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		reward := &rewardM.Reward{}

		selector := []interface{}{
			&reward.ID,
			&reward.CardRewardID,
			&reward.Order,
			&reward.Title,
			&reward.SubTitle,
			&reward.StartDate,
			&reward.EndDate,
			&reward.UpdateDate,
			&reward.PayloadOperator,
			&reward.Payloads,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)

			return nil, err
		}

		rewards = append(rewards, reward)

	}

	return rewards, nil
}

const UPDATE_BY_ID_STAT = "UPDATE reward SET " +
	" card_reward_id = $1, \"order\" = $2, \"title\" = $3, \"sub_title\" = $4, " +
	" start_date = $5, end_date = $6, update_date = $7, " +
	" payload_operator = $8, payload = $9 " +
	" WHERE \"id\" = $10"

func (im *impl) UpdateByID(ctx context.Context, reward *rewardM.Reward) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		reward.CardRewardID,
		reward.Order,
		reward.Title,
		reward.SubTitle,
		reward.StartDate,
		reward.EndDate,
		reward.UpdateDate,
		reward.PayloadOperator,
		reward.Payloads,
		reward.ID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		})
		return err
	}

	tx.Commit()
	return nil
}
