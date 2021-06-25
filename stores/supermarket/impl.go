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
	"(\"id\", \"name\", actionType, desc) VALUES ($1, $2, $3, $4)"

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
		supermarket.ActionType,
		supermarket.Desc,
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
	" \"name\" = $1, actionType = $2, desc = $3 " +
	" where \"id\" = $4"

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
		supermarket.ActionType,
		supermarket.Desc,
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

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", actionType, desc " +
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
			&supermarket.ActionType,
			&supermarket.Desc,
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