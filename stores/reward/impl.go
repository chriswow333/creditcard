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
	"(\"id\", \"card_id\", \"name\", \"desc\", start_date, end_date, update_date, constraint_payload)" +
	" VALUES($1, $2, $3, $4, $5, $6, $7, $8)"

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
		reward.CardID,
		reward.Name,
		reward.Desc,
		reward.StartDate,
		reward.EndDate,
		reward.UpdateDate,
		reward.ConstraintPayload,
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

const SELECT_STAT = "SELECT \"id\", card_id, \"name\", \"desc\", start_date, end_date, update_date, constraint_payload " +
	"FROM reward WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*rewardM.Reward, error) {

	reward := &rewardM.Reward{}

	selector := []interface{}{
		&reward.ID,
		&reward.CardID,
		&reward.Name,
		&reward.Desc,
		&reward.StartDate,
		&reward.EndDate,
		&reward.UpdateDate,
		&reward.ConstraintPayload,
	}

	if err := im.psql.QueryRow(SELECT_STAT, ID).Scan(selector...); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return reward, nil
}

const SELECT_BY_CARDID_STAT = "SELECT \"id\", card_id, \"name\", \"desc\", start_date, end_date, update_date, constraint_payload " +
	"FROM reward WHERE card_id = $1"

func (im *impl) GetByCardID(ctx context.Context, cardID string) ([]*rewardM.Reward, error) {

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
			&reward.CardID,
			&reward.Name,
			&reward.Desc,
			&reward.StartDate,
			&reward.EndDate,
			&reward.UpdateDate,
			&reward.ConstraintPayload,
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
	" card_id = $1, \"name\" = $2, \"desc\" = $3, start_date = $4, end_date = $5, update_date = $6, constraint_payload = $7 " +
	" WHERE \"id\" = $8"

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
		reward.CardID,
		reward.Name,
		reward.Desc,
		reward.StartDate,
		reward.EndDate,
		reward.UpdateDate,
		reward.ConstraintPayload,
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
