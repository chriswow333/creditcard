package onlinegame

import (
	"context"

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

const INSERT_STAT = "INSERT INTO onlinegame " +
	"(\"id\", \"name\", \"channel_label\") VALUES ($1, $2, $3)"

func (im *impl) Create(ctx context.Context, onlinegame *channel.Onlinegame) error {
	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		onlinegame.ID,
		onlinegame.Name,
		onlinegame.ChannelLabels,
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

const UPDATE_BY_ID_STAT = "UPDATE onlinegame SET " +
	" \"name\" = $1 " +
	" \"channel_label\" = $2 " +
	" where \"id\" = $3"

func (im *impl) UpdateByID(ctx context.Context, onlinegame *channel.Onlinegame) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		onlinegame.Name,
		onlinegame.ChannelLabels,
		onlinegame.ID,
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

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM onlinegame "

func (im *impl) GetAll(ctx context.Context) ([]*channel.Onlinegame, error) {
	onlinegames := []*channel.Onlinegame{}
	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		onlinegame := &channel.Onlinegame{}
		selector := []interface{}{
			&onlinegame.ID,
			&onlinegame.Name,
			&onlinegame.ChannelLabels,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		onlinegames = append(onlinegames, onlinegame)
	}

	return onlinegames, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", \"channel_label\" " +
	" FROM onlinegame WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*channel.Onlinegame, error) {
	onlinegame := &channel.Onlinegame{}

	selector := []interface{}{
		&onlinegame.ID,
		&onlinegame.Name,
		&onlinegame.ChannelLabels,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return onlinegame, nil
}
