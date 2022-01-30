package customization

import (
	"context"

	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	customizationM "example.com/creditcard/models/customization"
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

const INSERT_CUSTOMIZARION_STAT = "INSERT INTO customization " +
	"(\"id\", card_id, reward_id, \"name\", \"desc\", default_pass, link_url) VALUES ($1, $2, $3, $4, $5, $6, $7)"

func (im *impl) Create(ctx context.Context, customization *customizationM.Customization) error {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		customization.ID,
		customization.CardID,
		customization.RewardID,
		customization.Name,
		customization.Desc,
		customization.DefaultPass,
		customization.LinkURL,
	}

	if _, err := tx.Exec(INSERT_CUSTOMIZARION_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Fatal(err)

		return err
	}

	tx.Commit()

	return nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", card_id, reward_id, \"name\", \"desc\", default_pass, link_url " +
	"FROM customization " +
	" WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*customizationM.Customization, error) {

	customization := &customizationM.Customization{}

	selector := []interface{}{
		&customization.ID,
		&customization.CardID,
		&customization.RewardID,
		&customization.Name,
		&customization.Desc,
		&customization.DefaultPass,
		&customization.LinkURL,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err

	}
	return customization, nil
}

const UPDATE_BY_ID_STAT = "UPDATE customization SET " +
	" card_id = $1, reward_id = $2, \"name\" = $3, \"desc\" = $4, default_pass = $5, link_url = $6 " +
	" where \"id\" = $7"

func (im *impl) UpdateByID(ctx context.Context, customization *customizationM.Customization) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		customization.CardID,
		customization.RewardID,
		customization.Name,
		customization.Desc,
		customization.DefaultPass,
		customization.LinkURL,
		customization.ID,
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

const SELECT_BY_REWARD_ID_STAT = "SELECT \"id\", card_id, reward_id, \"name\", \"desc\", default_pass, link_url " +
	"FROM customization " +
	" WHERE reward_id = $1"

func (im *impl) GetByRewardID(ctx context.Context, rewardID string) ([]*customizationM.Customization, error) {
	customizations := []*customizationM.Customization{}

	rows, err := im.psql.Query(SELECT_BY_REWARD_ID_STAT, rewardID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		customization := &customizationM.Customization{}
		selector := []interface{}{
			&customization.ID,
			&customization.CardID,
			&customization.RewardID,
			&customization.Name,
			&customization.Desc,
			&customization.DefaultPass,
			&customization.LinkURL,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		customizations = append(customizations, customization)
	}
	return customizations, nil
}

const SELECT_BY_CARD_ID_STAT = "SELECT \"id\", card_id, reward_id, \"name\", \"desc\", default_pass, link_url " +
	"FROM customization " +
	" WHERE card_id = $1"

func (im *impl) GetByCardID(ctx context.Context, cardID string) ([]*customizationM.Customization, error) {
	customizations := []*customizationM.Customization{}

	rows, err := im.psql.Query(SELECT_BY_CARD_ID_STAT, cardID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		customization := &customizationM.Customization{}
		selector := []interface{}{
			&customization.ID,
			&customization.CardID,
			&customization.RewardID,
			&customization.Name,
			&customization.Desc,
			&customization.DefaultPass,
			&customization.LinkURL,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		customizations = append(customizations, customization)
	}
	return customizations, nil
}
