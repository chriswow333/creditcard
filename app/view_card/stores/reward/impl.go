package reward

import (
	"context"

	rewardM "example.com/creditcard/app/view_card/models/reward"
	"example.com/creditcard/app/view_card/utils/conn"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type impl struct {
	dig.In

	psql        *pgx.ConnPool `name:"psql"`
	connService conn.Service  `name:"connService"`
}

func New(
	psql *pgx.ConnPool,
	connService conn.Service,
) Store {
	return &impl{
		psql:        psql,
		connService: connService,
	}
}

const INSERT_STAT = "INSERT INTO reward " +
	" (\"id\", \"name\", card_id, \"desc\", reward_type, operator_type, " +
	" total_point, start_time, end_time,update_date) " +
	" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

func (im *impl) Create(ctx context.Context, conn *conn.Connection, reward *rewardM.Reward) error {

	updater := []interface{}{
		reward.ID,
		reward.Name,
		reward.CardID,
		reward.Desc,
		reward.RewardType,
		reward.OperatorType,
		reward.TotalPoint,
		reward.ValidateTime.StartTime,
		reward.ValidateTime.EndTime,
		reward.UpdateDate,
	}

	if err := im.connService.Exec(conn, INSERT_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Fatal(err)

		return err
	}

	return nil
}

const SELECT_STAT = "SELECT \"id\", \"name\", card_id, " +
	" \"desc\", reward_type, operator_type, total_point, " +
	" start_time, end_time, update_date " +
	" FROM reward " +
	" WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*rewardM.Reward, error) {
	reward := &rewardM.Reward{}

	selector := []interface{}{
		&reward.ID,
		&reward.Name,
		&reward.CardID,
		&reward.Desc,
		&reward.RewardType,
		&reward.OperatorType,
		&reward.TotalPoint,
		&reward.ValidateTime.StartTime,
		&reward.ValidateTime.EndTime,
		&reward.UpdateDate,
	}

	if err := im.psql.QueryRow(SELECT_STAT, ID).Scan(selector...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err

	}
	return reward, nil
}

const UPDATE_BY_ID_STAT = "UPDATE reward SET " +
	" \"name\" = $1, card_id = $2, \"desc\" = $3, " +
	" reward_type = $4, operator_type = $5, total_point = $6, " +
	" start_time = $7, end_time = $8, update_date = $9 " +
	" where \"id\" = $10"

func (im *impl) UpdateByID(ctx context.Context, conn *conn.Connection, reward *rewardM.Reward) error {

	updater := []interface{}{
		reward.Name,
		reward.CardID,
		reward.Desc,
		reward.RewardType,
		reward.OperatorType,
		reward.TotalPoint,
		reward.ValidateTime.StartTime,
		reward.ValidateTime.EndTime,
		reward.UpdateDate,
		reward.ID,
	}

	if err := im.connService.Exec(conn, UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		})
		return err
	}

	return nil
}

const SELECT_CARDID_STAT = "SELECT \"id\", \"name\", card_id, \"desc\", " +
	" reward_type, operator_type, total_point, start_time, end_time, update_date " +
	" FROM reward " +
	" WHERE card_id = $1"

func (im *impl) GetByCardID(ctx context.Context, cardID string) ([]*rewardM.Reward, error) {

	rewards := []*rewardM.Reward{}

	conditions := []interface{}{
		cardID,
	}
	rows, err := im.psql.Query(SELECT_CARDID_STAT, conditions...)

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
			&reward.Name,
			&reward.CardID,
			&reward.Desc,
			&reward.RewardType,
			&reward.OperatorType,
			&reward.TotalPoint,
			&reward.ValidateTime.StartTime,
			&reward.ValidateTime.EndTime,
			&reward.UpdateDate,
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
