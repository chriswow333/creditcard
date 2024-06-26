package reward

import (
	"context"
	"runtime/debug"

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
	"(\"id\", \"card_reward_id\", \"order\", payload_operator, payload) " +
	" VALUES($1, $2, $3, $4, $5)"

func (im *impl) Create(ctx context.Context, reward *rewardM.Reward) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		reward.ID,
		reward.CardRewardID,
		reward.Order,
		reward.PayloadOperator,
		reward.Payloads,
	}
	if _, err := tx.Exec(INSERT_REWARD_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_STAT = "SELECT \"id\", card_reward_id, \"order\", payload_operator, payload " +
	" FROM reward WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*rewardM.Reward, error) {

	reward := &rewardM.Reward{}

	selector := []interface{}{
		&reward.ID,
		&reward.CardRewardID,
		&reward.Order,
		&reward.PayloadOperator,
		&reward.Payloads,
	}

	if err := im.psql.QueryRow(SELECT_STAT, ID).Scan(selector...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return reward, nil
}

const SELECT_BY_CARDID_STAT = "SELECT \"id\", card_reward_id, \"order\", payload_operator, payload " +
	" FROM reward WHERE card_reward_id = $1"

func (im *impl) GetByCardRewardID(ctx context.Context, cardID string) ([]*rewardM.Reward, error) {

	rewards := []*rewardM.Reward{}

	conditions := []interface{}{
		cardID,
	}
	rows, err := im.psql.Query(SELECT_BY_CARDID_STAT, conditions...)

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		reward := &rewardM.Reward{}

		selector := []interface{}{
			&reward.ID,
			&reward.CardRewardID,
			&reward.Order,
			&reward.PayloadOperator,
			&reward.Payloads,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return nil, err
		}

		rewards = append(rewards, reward)

	}

	return rewards, nil
}

const UPDATE_BY_ID_STAT = "UPDATE reward SET " +
	" card_reward_id = $1, \"order\" = $2, payload_operator = $3, payload = $4 " +
	" WHERE \"id\" = $5"

func (im *impl) UpdateByID(ctx context.Context, reward *rewardM.Reward) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		reward.CardRewardID,
		reward.Order,
		reward.PayloadOperator,
		reward.Payloads,
		reward.ID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()
	return nil
}
