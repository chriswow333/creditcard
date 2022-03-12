package card_reward

import (
	"context"

	cardM "example.com/creditcard/models/card"
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

const INSERT_CARD_REWARD_STAT = "INSERT INTO card_reward " +
	"(\"id\", \"card_id\", \"reward_operator\", \"reward_type\") " +
	" VALUES($1, $2, $3, $4)"

func (im *impl) Create(ctx context.Context, cardReward *cardM.CardReward) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		cardReward.ID,
		cardReward.CardID,
		cardReward.RewardOperator,
		cardReward.RewardType,
	}
	if _, err := tx.Exec(INSERT_CARD_REWARD_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		})
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_BY_CARDID_STAT = "SELECT \"id\", card_id, reward_operator, reward_type " +
	" FROM card_reward " +
	" WHERE \"card_id\"=$1"

func (im *impl) GetByCardID(ctx context.Context, cardID string) ([]*cardM.CardReward, error) {
	cardRewards := []*cardM.CardReward{}

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

		cardReward := &cardM.CardReward{}
		selector := []interface{}{
			&cardReward.ID,
			&cardReward.CardID,
			&cardReward.RewardOperator,
			&cardReward.RewardType,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		cardRewards = append(cardRewards, cardReward)
	}

	return cardRewards, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", card_id, reward_operator, reward_type " +
	" FROM card_reward " +
	" WHERE \"id\"=$1"

func (im *impl) GetByID(ctx context.Context, ID string) (*cardM.CardReward, error) {
	cardReward := &cardM.CardReward{}

	selector := []interface{}{
		&cardReward.ID,
		&cardReward.CardID,
		&cardReward.RewardOperator,
		&cardReward.RewardType,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return cardReward, nil
}
