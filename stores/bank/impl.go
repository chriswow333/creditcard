package bank

import (
	"context"

	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"

	bankM "example.com/creditcard/models/bank"
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

const INSERT_BANK_STAT = "INSERT INTO bank " +
	" (\"id\", \"name\", update_date, image_path, link_url) " +
	" VALUES($1, $2, $3, $4, $5)"

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
		bank.UpdateDate,
		bank.ImagePath,
		bank.LinkURL,
	}

	if _, err := tx.Exec(INSERT_BANK_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Fatal(err)

		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE bank SET " +
	" \"name\" = $1, update_date = $2, image_path = $3, link_url = $4 " +
	" where \"id\" = $5"

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
		bank.UpdateDate,
		bank.ImagePath,
		bank.LinkURL,
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

const SELECT_STAT = "SELECT \"id\", \"name\", update_date, image_path, link_url " +
	" FROM bank " +
	" WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*bankM.Bank, error) {

	bank := &bankM.Bank{}

	selector := []interface{}{
		&bank.ID,
		&bank.Name,
		&bank.UpdateDate,
		&bank.ImagePath,
		&bank.LinkURL,
	}

	if err := im.psql.QueryRow(SELECT_STAT, ID).Scan(selector...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err

	}

	return bank, nil
}

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", update_date, image_path, link_url " +
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
			&bank.UpdateDate,
			&bank.ImagePath,
			&bank.LinkURL,
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
