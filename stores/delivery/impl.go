package delivery

import (
	"context"

	"example.com/creditcard/models/delivery"
	delivertM "example.com/creditcard/models/delivery"
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

const INSERT_DELIVERY_STAT = "INSERT INTO delivery " +
	"(\"id\", \"name\", \"desc\", link_url) VALUES ($1, $2, $3, $4)"

func (im *impl) Create(ctx context.Context, delivery *delivertM.Delivery) error {
	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		delivery.ID,
		delivery.Name,
		delivery.Desc,
		delivery.LinkURL,
	}

	if _, err := tx.Exec(INSERT_DELIVERY_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Fatal(err)

		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE delivery SET " +
	" \"name\" = $1, \"desc\" = $2, link_url = $3 " +
	" where \"id\" = $4"

func (im *impl) UpdateByID(ctx context.Context, delivery *delivery.Delivery) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		delivery.Name,
		delivery.Desc,
		delivery.LinkURL,
		delivery.ID,
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

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"desc\", link_url " +
	" FROM delivery "

func (im *impl) GetAll(ctx context.Context) ([]*delivertM.Delivery, error) {
	deliverys := []*delivertM.Delivery{}
	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		delivery := &delivertM.Delivery{}
		selector := []interface{}{
			&delivery.ID,
			&delivery.Name,
			&delivery.Desc,
			&delivery.LinkURL,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		deliverys = append(deliverys, delivery)
	}

	return deliverys, nil
}
