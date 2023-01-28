package appstore

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

const INSERT_STAT = "INSERT INTO appstore " +
	"(\"id\", \"name\", \"channel_label\") VALUES ($1, $2, $3)"

func (im *impl) Create(ctx context.Context, appstore *channel.AppStore) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		appstore.ID,
		appstore.Name,
		appstore.ChannelLabels,
	}

	if _, err := tx.Exec(INSERT_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE appstore SET " +
	" \"name\" = $1 " +
	" \"channel_label\" = $2 " +
	" where \"id\" = $3"

func (im *impl) UpdateByID(ctx context.Context, appstore *channel.AppStore) error {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		appstore.Name,
		appstore.ChannelLabels,
		appstore.ID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM appstore "

func (im *impl) GetAll(ctx context.Context) ([]*channel.AppStore, error) {

	appstores := []*channel.AppStore{}

	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		appstore := &channel.AppStore{}
		selector := []interface{}{
			&appstore.ID,
			&appstore.Name,
			&appstore.ChannelLabels,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
			return nil, err
		}

		appstores = append(appstores, appstore)
	}

	return appstores, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM appstore WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*channel.AppStore, error) {
	appstore := &channel.AppStore{}

	selector := []interface{}{
		&appstore.ID,
		&appstore.Name,
		&appstore.ChannelLabels,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	return appstore, nil
}

const SELECT_BY_LIKE_NAME_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM appstore WHERE \"name\" ~* $1"

func (im *impl) FindLike(ctx context.Context, names []string) ([]*channel.AppStore, error) {
	appstores := []*channel.AppStore{}

	name := strings.Join(names, "|")

	rows, err := im.psql.Query(SELECT_BY_LIKE_NAME_STAT, name)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		appstore := &channel.AppStore{}
		selector := []interface{}{
			&appstore.ID,
			&appstore.Name,
			&appstore.ChannelLabels,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
			return nil, err
		}

		appstores = append(appstores, appstore)
	}

	return appstores, nil
}
