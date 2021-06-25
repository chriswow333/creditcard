package customization

import (
	"context"

	customizationM "example.com/creditcard/models/customization"
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

const INSERT_CUSTOMIZARION_STAT = "INSERT INTO customization " +
	"(\"id\", reward_id, \"name\", descs) VALUES ($1, $2, $3, $4)"

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
		customization.RewardID,
		customization.Name,
		customization.Descs,
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

const SELECT_BY_ID_STAT = "SELECT \"id\", reward_id, \"name\", descs " +
	"FROM customization " +
	" WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*customizationM.Customization, error) {

	customization := &customizationM.Customization{}

	selector := []interface{}{
		&customization.ID,
		&customization.RewardID,
		&customization.Name,
		&customization.Descs,
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
	" reward_id = $1, \"name\" = $2, descs = $3 " +
	" where \"id\" = $4"

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
		customization.RewardID,
		customization.Name,
		customization.Descs,
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

const SELECT_BY_REWARDID_STAT = "SELECT \"id\", reward_id, \"name\", descs " +
	" FROM customization " +
	" WHERE reward_id = $1"

func (im *impl) GetByRewardID(ctx context.Context, rewardID string) ([]*customizationM.Customization, error) {

	customizations := []*customizationM.Customization{}
	rows, err := im.psql.Query(SELECT_BY_REWARDID_STAT)
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
			&customization.RewardID,
			&customization.Name,
			&customization.Descs,
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
