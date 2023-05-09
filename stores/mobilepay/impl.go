package mobilepay

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

const INSERT_STAT = "INSERT INTO mobilepay " +
	"(\"id\", \"name\", \"channel_label\", \"image_path\") VALUES ($1, $2, $3)"

func (im *impl) Create(ctx context.Context, mobilepay *channel.Mobilepay) error {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		mobilepay.ID,
		mobilepay.Name,
		mobilepay.ChannelLabels,
		mobilepay.ImagePath,
	}

	if _, err := tx.Exec(INSERT_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE mobilepay SET " +
	" \"name\" = $1 " +
	" \"channel_label\" = $2 " +
	" \"image_path\" = $3 " +
	" where \"id\" = $4"

func (im *impl) UpdateByID(ctx context.Context, mobilepay *channel.Mobilepay) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		mobilepay.Name,
		mobilepay.ChannelLabels,
		mobilepay.ImagePath,
		mobilepay.ID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"channel_label\", \"image_path\" " +
	" FROM mobilepay "

func (im *impl) GetAll(ctx context.Context) ([]*channel.Mobilepay, error) {

	mobilepays := []*channel.Mobilepay{}

	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		mobilepay := &channel.Mobilepay{}
		selector := []interface{}{
			&mobilepay.ID,
			&mobilepay.Name,
			&mobilepay.ChannelLabels,
			&mobilepay.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return nil, err
		}

		mobilepays = append(mobilepays, mobilepay)
	}

	return mobilepays, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", \"channel_label\", \"image_path\" " +
	" FROM mobilepay WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*channel.Mobilepay, error) {

	mobilepay := &channel.Mobilepay{}

	selector := []interface{}{
		&mobilepay.ID,
		&mobilepay.Name,
		&mobilepay.ChannelLabels,
		&mobilepay.ImagePath,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return mobilepay, nil
}

const SELECT_BY_LIKE_NAME_STAT = "SELECT \"id\", \"name\", \"channel_label\", \"image_path\" " +
	" FROM mobilepay WHERE \"name\" ~* $1"

func (im *impl) FindLike(ctx context.Context, names []string) ([]*channel.Mobilepay, error) {
	mobilepays := []*channel.Mobilepay{}

	name := strings.Join(names, "|")

	rows, err := im.psql.Query(SELECT_BY_LIKE_NAME_STAT, name)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		mobilepay := &channel.Mobilepay{}
		selector := []interface{}{
			&mobilepay.ID,
			&mobilepay.Name,
			&mobilepay.ChannelLabels,
			&mobilepay.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
			return nil, err
		}

		mobilepays = append(mobilepays, mobilepay)
	}

	return mobilepays, nil
}
