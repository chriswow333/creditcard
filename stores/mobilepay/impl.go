package mobilepay

import (
	"context"

	"example.com/creditcard/models/mobilepay"
	mobilepayM "example.com/creditcard/models/mobilepay"
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

const INSERT_STAT = "INSERT INTO mobilepay " +
	"(\"id\", \"name\", desc, link_url) VALUES ($1, $2, $3, $4)"

func (im *impl) Create(ctx context.Context, mobilepay *mobilepayM.Mobilepay) error {
	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		mobilepay.ID,
		mobilepay.Name,
		mobilepay.Desc,
		mobilepay.LinkURL,
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

const UPDATE_BY_ID_STAT = "UPDATE mobilepay SET " +
	" \"name\" = $1, desc = $2, link_url = $3 " +
	" where \"id\" = $4"

func (im *impl) UpdateByID(ctx context.Context, mobilepay *mobilepay.Mobilepay) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		mobilepay.Name,
		mobilepay.Desc,
		mobilepay.LinkURL,
		mobilepay.ID,
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

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", desc, link_url " +
	" FROM mobilepay "

func (im *impl) GetAll(ctx context.Context) ([]*mobilepayM.Mobilepay, error) {
	mobilepays := []*mobilepayM.Mobilepay{}
	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		mobilepay := &mobilepayM.Mobilepay{}
		selector := []interface{}{
			&mobilepay.ID,
			&mobilepay.Name,
			&mobilepay.Desc,
			&mobilepay.LinkURL,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		mobilepays = append(mobilepays, mobilepay)
	}

	return mobilepays, nil
}
