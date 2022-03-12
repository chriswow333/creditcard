package streaming

import (
	"context"

	streamingM "example.com/creditcard/models/streaming"
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
	"(\"id\", \"name\", \"image_path\") VALUES ($1, $2, $3)"

func (im *impl) Create(ctx context.Context, streaming *streamingM.Streaming) error {
	tx, err := im.psql.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"msg": "",
		}).Error(err)
		return err
	}

	defer tx.Rollback()

	updater := []interface{}{
		streaming.ID,
		streaming.Name,
		streaming.ImagePath,
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

const UPDATE_BY_ID_STAT = "UPDATE streaming SET " +
	" \"name\" = $1, \"image_path\" = $2 " +
	" where \"id\" = $3"

func (im *impl) UpdateByID(ctx context.Context, streaming *streamingM.Streaming) error {
	tx, err := im.psql.Begin()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return err
	}
	defer tx.Rollback()

	updater := []interface{}{
		streaming.Name,
		streaming.ImagePath,
		streaming.ID,
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
	" FROM streaming "

func (im *impl) GetAll(ctx context.Context) ([]*streamingM.Streaming, error) {
	streamings := []*streamingM.Streaming{}
	rows, err := im.psql.Query(SELECT_ALL_STAT)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"": "",
		}).Error(err)
		return nil, err
	}

	for rows.Next() {

		streaming := &streamingM.Streaming{}
		selector := []interface{}{
			&streaming.ID,
			&streaming.Name,
			&streaming.ImagePath,
		}

		if err := rows.Scan(selector...); err != nil {
			logrus.WithFields(logrus.Fields{
				"": "",
			}).Error(err)
			return nil, err
		}

		streamings = append(streamings, streaming)
	}

	return streamings, nil
}
