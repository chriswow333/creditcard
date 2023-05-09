package supermarket

import (
	"context"
	"runtime/debug"
	"strings"

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

const INSERT_STAT = "INSERT INTO supermarket " +
	"(\"id\", \"name\", \"channel_label\", \"image_path\") VALUES ($1, $2, $3, $4)"

func (im *impl) Create(ctx context.Context, supermarket *channel.Supermarket) error {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		supermarket.ID,
		supermarket.Name,
		supermarket.ChannelLabels,
		supermarket.ImagePath,
	}

	if _, err := tx.Exec(INSERT_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE supermarket SET " +
	" \"name\" = $1 " +
	" \"channel_label\" = $2 " +
	" \"image_path\" = $3 " +
	" where \"id\" = $4"

func (im *impl) UpdateByID(ctx context.Context, supermarket *channel.Supermarket) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		supermarket.Name,
		supermarket.ChannelLabels,
		supermarket.ImagePath,
		supermarket.ID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"channel_label\", \"image_path\" " +
	" FROM supermarket "

func (im *impl) GetAll(ctx context.Context) ([]*channel.Supermarket, error) {

	supermarkets := []*channel.Supermarket{}

	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		supermarket := &channel.Supermarket{}
		selector := []interface{}{
			&supermarket.ID,
			&supermarket.Name,
			&supermarket.ChannelLabels,
			&supermarket.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return nil, err
		}

		supermarkets = append(supermarkets, supermarket)
	}

	return supermarkets, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", \"channel_label\", \"image_path\" " +
	" FROM supermarket WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*channel.Supermarket, error) {

	supermarket := &channel.Supermarket{}

	selector := []interface{}{
		&supermarket.ID,
		&supermarket.Name,
		&supermarket.ChannelLabels,
		&supermarket.ImagePath,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return supermarket, nil
}

const SELECT_BY_LIKE_NAME_STAT = "SELECT \"id\", \"name\", \"channel_label\", \"image_path\" " +
	" FROM supermarket WHERE \"name\" ~* $1"

func (im *impl) FindLike(ctx context.Context, names []string) ([]*channel.Supermarket, error) {
	supermarkets := []*channel.Supermarket{}

	name := strings.Join(names, "|")

	rows, err := im.psql.Query(SELECT_BY_LIKE_NAME_STAT, name)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		supermarket := &channel.Supermarket{}
		selector := []interface{}{
			&supermarket.ID,
			&supermarket.Name,
			&supermarket.ChannelLabels,
			&supermarket.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
			return nil, err
		}

		supermarkets = append(supermarkets, supermarket)
	}

	return supermarkets, nil
}
