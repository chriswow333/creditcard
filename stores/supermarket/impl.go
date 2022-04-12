package supermarket

import (
	"context"

	supermarketM "example.com/creditcard/models/supermarket"
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

const INSERT_STAT = "INSERT INTO supermarket " +
	"(\"id\", \"name\", \"image_path\") VALUES ($1, $2, $3)"

func (im *impl) Create(ctx context.Context, supermarket *supermarketM.Supermarket) error {
	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		supermarket.ID,
		supermarket.Name,
		supermarket.ImagePath,
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

const UPDATE_BY_ID_STAT = "UPDATE supermarket SET " +
	" \"name\" = $1, \"image_path\" = $2 " +
	" where \"id\" = $3"

func (im *impl) UpdateByID(ctx context.Context, supermarket *supermarketM.Supermarket) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		supermarket.Name,
		supermarket.ImagePath,
		supermarket.ID,
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
	" FROM supermarket "

func (im *impl) GetAll(ctx context.Context) ([]*supermarketM.Supermarket, error) {
	supermarkets := []*supermarketM.Supermarket{}
	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		supermarket := &supermarketM.Supermarket{}
		selector := []interface{}{
			&supermarket.ID,
			&supermarket.Name,
			&supermarket.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		supermarkets = append(supermarkets, supermarket)
	}

	return supermarkets, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", \"image_path\" " +
	" FROM supermarket WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*supermarketM.Supermarket, error) {
	supermarket := &supermarketM.Supermarket{}

	selector := []interface{}{
		&supermarket.ID,
		&supermarket.Name,
		&supermarket.ImagePath,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return supermarket, nil
}
