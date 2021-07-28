package bank

import (
	"context"

	bankM "example.com/creditcard/app/view_card/models/bank"
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

const INSERT_STAT = "INSERT INTO bank " +
	" (\"id\", \"name\", icon, update_date) " +
	" VALUES ($1, $2, $3, $4)"

func (im *impl) Create(ctx context.Context, bank *bankM.Bank) error {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		bank.ID,
		bank.Name,
		bank.Icon,
		bank.UpdateDate,
	}

	if _, err := tx.Exec(INSERT_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Fatal(err)

		return err
	}

	tx.Commit()

	return nil
}

const SELECT_STAT = "SELECT \"id\", \"name\", icon, update_date " +
	" FROM bank " +
	" WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*bankM.Bank, error) {

	bank := &bankM.Bank{}

	selector := []interface{}{
		&bank.ID,
		&bank.Name,
		&bank.Icon,
		&bank.UpdateDate,
	}

	if err := im.psql.QueryRow(SELECT_STAT, ID).Scan(selector...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err

	}
	return bank, nil
}

const UPDATE_BY_ID_STAT = "UPDATE bank SET " +
	" \"name\" = $1, icon = $2, update_date = $3 " +
	" where \"id\" = $4"

func (im *impl) UpdateByID(ctx context.Context, bank *bankM.Bank) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		bank.Name,
		bank.Icon,
		bank.UpdateDate,
		bank.ID,
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

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", icon, update_date " +
	" FROM bank"

func (im *impl) GetAll(ctx context.Context) ([]*bankM.Bank, error) {

	banks := []*bankM.Bank{}
	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		bank := &bankM.Bank{}
		selector := []interface{}{
			&bank.ID,
			&bank.Name,
			&bank.Icon,
			&bank.UpdateDate,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		banks = append(banks, bank)
	}

	return banks, nil
}
