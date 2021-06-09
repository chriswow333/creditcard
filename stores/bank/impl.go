package bank

import (
	"context"

	uuid "github.com/nu7hatch/gouuid"

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

const INSERT_BANK_STAT = "INSERT INTO bank(\"id\", \"name\", \"desc\", start_date, end_date, update_date) VALUES($1, $2, $3, $4, $5, $6)"

func (im *impl) Create(ctx context.Context, bank *bankM.Bank) error {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	id, err := uuid.NewV4()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Fatal(err)
		return err
	}

	updater := []interface{}{
		id.String(),
		bank.Name,
		bank.Desc,
		bank.StartDate,
		bank.EndDate,
		bank.UpdateDate,
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

const SELECT_STAT = "SELECT \"id\", \"name\", \"desc\", start_date, end_date, update_date FROM bank WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*bankM.Bank, error) {

	bank := &bankM.Bank{}

	selector := []interface{}{
		&bank.ID,
		&bank.Name,
		&bank.Desc,
		&bank.StartDate,
		&bank.EndDate,
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

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"desc\", start_date, end_date, update_date FROM bank"

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
			&bank.Desc,
			&bank.StartDate,
			&bank.EndDate,
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
