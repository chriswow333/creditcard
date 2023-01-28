package amusement

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

const INSERT_STAT = "INSERT INTO amusement " +
	"(\"id\", \"name\", \"channel_label\") VALUES ($1, $2, $3)"

func (im *impl) Create(ctx context.Context, amusement *channel.Amusement) error {

	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		amusement.ID,
		amusement.Name,
		amusement.ChannelLabels,
	}

	if _, err := tx.Exec(INSERT_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE amusement SET " +
	" \"name\" = $1 " +
	" \"channel_label\" = $2 " +
	" where \"id\" = $3"

func (im *impl) UpdateByID(ctx context.Context, amusement *channel.Amusement) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		amusement.Name,
		amusement.ChannelLabels,
		amusement.ID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM amusement "

func (im *impl) GetAll(ctx context.Context) ([]*channel.Amusement, error) {
	amusements := []*channel.Amusement{}
	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		amusement := &channel.Amusement{}
		selector := []interface{}{
			&amusement.ID,
			&amusement.Name,
			&amusement.ChannelLabels,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return nil, err
		}

		amusements = append(amusements, amusement)
	}

	return amusements, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM amusement WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*channel.Amusement, error) {
	amusement := &channel.Amusement{}

	selector := []interface{}{
		&amusement.ID,
		&amusement.Name,
		&amusement.ChannelLabels,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return amusement, nil
}

const SELECT_BY_LIKE_NAME_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM amusement WHERE \"name\" ~* $1"

func (im *impl) FindLike(ctx context.Context, names []string) ([]*channel.Amusement, error) {
	amusements := []*channel.Amusement{}

	name := strings.Join(names, "|")
	logrus.Info(name)
	rows, err := im.psql.Query(SELECT_BY_LIKE_NAME_STAT, name)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		amusement := &channel.Amusement{}
		selector := []interface{}{
			&amusement.ID,
			&amusement.Name,
			&amusement.ChannelLabels,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
			return nil, err
		}

		amusements = append(amusements, amusement)
	}

	return amusements, nil
}
