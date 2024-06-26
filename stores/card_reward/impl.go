package card_reward

import (
	"context"
	"runtime/debug"

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
	"(\"id\", \"card_id\", \"card_reward_operator\", \"title\", \"descs\", start_date, end_date, \"reward_type\", " +
	" constraint_pass_logic, card_reward_limit_types, feedback_bonus) " +
	" VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"

func (im *impl) Create(ctx context.Context, cardReward *cardM.CardReward) error {

	tx, err := im.psql.Begin()
	defer tx.Rollback()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	updater := []interface{}{
		cardReward.ID,
		cardReward.CardID,
		cardReward.CardRewardOperator,
		cardReward.Title,
		cardReward.Descs,
		cardReward.StartDate,
		cardReward.EndDate,
		cardReward.RewardType,
		cardReward.ConstraintPassLogics,
		cardReward.CardRewardLimitTypes,
		cardReward.FeedbackBonus,
	}
	if _, err := tx.Exec(INSERT_CARD_REWARD_STAT, updater...); err != nil {
		logrus.Error(err)
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_BY_CARDID_STAT = "SELECT \"id\", card_id, card_reward_operator, \"title\", \"descs\", " +
	" start_date, end_date, reward_type, constraint_pass_logic, card_reward_limit_types, feedback_bonus " +
	" FROM card_reward " +
	" WHERE \"card_id\"=$1"

func (im *impl) GetByCardID(ctx context.Context, cardID string) ([]*cardM.CardReward, error) {
	cardRewards := []*cardM.CardReward{}

	conditions := []interface{}{
		cardID,
	}
	rows, err := im.psql.Query(SELECT_BY_CARDID_STAT, conditions...)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		cardReward := &cardM.CardReward{}
		selector := []interface{}{
			&cardReward.ID,
			&cardReward.CardID,
			&cardReward.CardRewardOperator,
			&cardReward.Title,
			&cardReward.Descs,
			&cardReward.StartDate,
			&cardReward.EndDate,
			&cardReward.RewardType,
			&cardReward.ConstraintPassLogics,
			&cardReward.CardRewardLimitTypes,
			&cardReward.FeedbackBonus,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return nil, err
		}

		cardRewards = append(cardRewards, cardReward)
	}

	return cardRewards, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", card_id, card_reward_operator, \"title\", \"descs\", " +
	" start_date, end_date, reward_type, constraint_pass_logic, card_reward_limit_types, feedback_bonus " +
	" FROM card_reward " +
	" WHERE \"id\"=$1"

func (im *impl) GetByID(ctx context.Context, ID string) (*cardM.CardReward, error) {
	cardReward := &cardM.CardReward{}

	selector := []interface{}{
		&cardReward.ID,
		&cardReward.CardID,
		&cardReward.CardRewardOperator,
		&cardReward.Title,
		&cardReward.Descs,
		&cardReward.StartDate,
		&cardReward.EndDate,
		&cardReward.RewardType,
		&cardReward.ConstraintPassLogics,
		&cardReward.CardRewardLimitTypes,
		&cardReward.FeedbackBonus,
		// &cardReward.FeedbackDescID,
		// &cardReward.CardRewardBonus,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return cardReward, nil
}
