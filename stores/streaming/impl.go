package streaming

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

const INSERT_STAT = "INSERT INTO streaming " +
	"(\"id\", \"name\", \"channel_label\", \"image_path\") VALUES ($1, $2, $3, $4)"

func (im *impl) Create(ctx context.Context, streaming *channel.Streaming) error {

	tx, err := im.psql.Begin()
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		streaming.ID,
		streaming.Name,
		streaming.ChannelLabels,
		streaming.ImagePath,
	}

	if _, err := tx.Exec(INSERT_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()

	return nil
}

const UPDATE_BY_ID_STAT = "UPDATE streaming SET " +
	" \"name\" = $1 " +
	" \"channel_label\" = $2 " +
	" \"image_path\" = $3 " +
	" where \"id\" = $4"

func (im *impl) UpdateByID(ctx context.Context, streaming *channel.Streaming) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		streaming.Name,
		streaming.ChannelLabels,
		streaming.ImagePath,
		streaming.ID,
	}

	if _, err := tx.Exec(UPDATE_BY_ID_STAT, updater...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return err
	}

	tx.Commit()
	return nil
}

const SELECT_ALL_STAT = "SELECT \"id\", \"name\", \"channel_label\", \"image_path\" " +
	" FROM streaming limit $1 offset $2 "

func (im *impl) GetAll(ctx context.Context, offset, limit int) ([]*channel.Streaming, error) {

	streamings := []*channel.Streaming{}

	rows, err := im.psql.Query(SELECT_ALL_STAT, limit, offset)
	if err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		streaming := &channel.Streaming{}
		selector := []interface{}{
			&streaming.ID,
			&streaming.Name,
			&streaming.ChannelLabels,
			&streaming.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
			return nil, err
		}

		streamings = append(streamings, streaming)
	}

	return streamings, nil
}

const SELECT_BY_ID_STAT = "SELECT \"id\", \"name\", \"channel_label\", \"image_path\" " +
	" FROM streaming WHERE \"id\" = $1"

func (im *impl) GetByID(ctx context.Context, ID string) (*channel.Streaming, error) {

	streaming := &channel.Streaming{}

	selector := []interface{}{
		&streaming.ID,
		&streaming.Name,
		&streaming.ChannelLabels,
		&streaming.ImagePath,
	}

	if err := im.psql.QueryRow(SELECT_BY_ID_STAT, ID).Scan(selector...); err != nil {
		logrus.Errorf("[PANIC] \n%s", string(debug.Stack()))
		return nil, err
	}

	return streaming, nil
}

const SELECT_BY_LIKE_NAME_STAT = "SELECT \"id\", \"name\", \"channel_label\", \"image_path\" " +
	" FROM streaming WHERE \"name\" ~* $1"

func (im *impl) FindLike(ctx context.Context, names []string) ([]*channel.Streaming, error) {
	streamings := []*channel.Streaming{}

	name := strings.Join(names, "|")

	rows, err := im.psql.Query(SELECT_BY_LIKE_NAME_STAT, name)
	if err != nil {
		logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
		return nil, err
	}

	for rows.Next() {

		streaming := &channel.Streaming{}
		selector := []interface{}{
			&streaming.ID,
			&streaming.Name,
			&streaming.ChannelLabels,
			&streaming.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.Errorf("[PANIC] %s\n%s", err, string(debug.Stack()))
			return nil, err
		}

		streamings = append(streamings, streaming)
	}

	return streamings, nil
}
