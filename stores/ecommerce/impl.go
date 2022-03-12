package ecommerce

import (
	"context"

	"example.com/creditcard/models/ecommerce"
	ecommerceM "example.com/creditcard/models/ecommerce"
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

const INSERT_ECOMMERCE_STAT = "INSERT INTO ecommerce " +
	"(\"id\", \"name\", \"image_path\") VALUES ($1, $2, $3)"

func (im *impl) Create(ctx context.Context, ecommerce *ecommerceM.Ecommerce) error {
	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		ecommerce.ID,
		ecommerce.Name,
		ecommerce.ImagePath,
	}

	if _, err := tx.Exec(INSERT_ECOMMERCE_STAT, updater...); err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Fatal(err)

		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE ecommerce SET " +
	" \"name\" = $1, image_path = $2 " +
	" where \"id\" = $3"

func (im *impl) UpdateByID(ctx context.Context, ecommerce *ecommerce.Ecommerce) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		ecommerce.Name,
		ecommerce.ImagePath,
		ecommerce.ID,
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

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"image_path\" " +
	" FROM ecommerce "

func (im *impl) GetAll(ctx context.Context) ([]*ecommerceM.Ecommerce, error) {
	ecommerces := []*ecommerceM.Ecommerce{}
	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		ecommerce := &ecommerceM.Ecommerce{}
		selector := []interface{}{
			&ecommerce.ID,
			&ecommerce.Name,
			&ecommerce.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		ecommerces = append(ecommerces, ecommerce)
	}

	return ecommerces, nil
}
