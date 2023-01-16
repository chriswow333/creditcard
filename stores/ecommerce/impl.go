package ecommerce

import (
	"context"
	"runtime/debug"

	"example.com/creditcard/models/channel"
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
	"(\"id\", \"name\", \"channel_label\") VALUES ($1, $2, $3)"

func (im *impl) Create(ctx context.Context, ecommerce *channel.Ecommerce) error {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		ecommerce.ID,
		ecommerce.Name,
		ecommerce.ChannelLabels,
	}

	if _, err := tx.Exec(INSERT_ECOMMERCE_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE ecommerce SET " +
	" \"name\" = $1 " +
	" \"channel_label\" = $2  " +
	" where \"id\" = $3"

func (im *impl) UpdateByID(ctx context.Context, ecommerce *channel.Ecommerce) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		ecommerce.Name,
		ecommerce.ChannelLabels,
		ecommerce.ID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM ecommerce "

func (im *impl) GetAll(ctx context.Context) ([]*channel.Ecommerce, error) {
	ecommerces := []*channel.Ecommerce{}
	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		ecommerce := &channel.Ecommerce{}
		selector := []interface{}{
			&ecommerce.ID,
			&ecommerce.Name,
			&ecommerce.ChannelLabels,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return nil, err
		}

		ecommerces = append(ecommerces, ecommerce)
	}

	return ecommerces, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM ecommerce WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*channel.Ecommerce, error) {

	ecommerce := &channel.Ecommerce{}

	selector := []interface{}{
		&ecommerce.ID,
		&ecommerce.Name,
		&ecommerce.ChannelLabels,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return ecommerce, nil
}
